terraform {
  required_providers {
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 5.0"
    }
  }

  backend "gcs" {
    bucket = "stredono-terraform-state"
    prefix = "terraform/state"
  }
}

locals {
  base_path  = "${path.module}/.."
  rules_path = "${local.base_path}/rules"
}

provider "google-beta" {
  user_project_override = true
}

provider "google-beta" {
  alias                 = "no_user_project_override"
  user_project_override = false
}

resource "random_id" "project-id" {
  byte_length = 4
}

resource "google_project" "default" {
  provider   = google-beta.no_user_project_override
  name       = "Stredono"
  project_id = "stredono-${random_id.project-id.hex}"

  billing_account = "0152B3-F5E957-35889E"

  labels = {
    "firebase" = "enabled"
  }
}

resource "google_project_service" "default" {
  provider = google-beta.no_user_project_override
  project  = google_project.default.project_id
  for_each = toset([
    "cloudbilling.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "serviceusage.googleapis.com",
    "identitytoolkit.googleapis.com",
    "firebase.googleapis.com",
    "firebaserules.googleapis.com",
    "firestore.googleapis.com",
    "pubsub.googleapis.com",
    "cloudfunctions.googleapis.com",
    "run.googleapis.com",
    "cloudbuild.googleapis.com",
    "storage.googleapis.com",
    "firebasestorage.googleapis.com",
    "firebasedatabase.googleapis.com",
    "firebaseappcheck.googleapis.com",
    "secretmanager.googleapis.com",
    "recaptchaenterprise.googleapis.com",
    "cloudkms.googleapis.com",
    "iam.googleapis.com",
  ])

  service            = each.key
  disable_on_destroy = false
}


resource "google_project_iam_member" "default" {
  provider   = google-beta
  project    = google_project.default.project_id
  role       = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member     = "serviceAccount:service-${google_project.default.number}@gs-project-accounts.iam.gserviceaccount.com"
  depends_on = [google_project_service.default]
}

resource "google_kms_key_ring" "default" {
  provider = google-beta
  project  = google_project.default.project_id

  location   = var.kms_location
  name       = "${google_project.default.project_id}-keyring"
  depends_on = [google_project_iam_member.default]
}

resource "google_kms_crypto_key" "default" {
  name     = "${google_project.default.project_id}-key"
  key_ring = google_kms_key_ring.default.id

  rotation_period = "86400s"

  lifecycle {
    prevent_destroy = false
  }

  depends_on = [google_kms_key_ring.default]
}

resource "google_storage_bucket" "terraform_state" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "stredono-terraform-state"
  location = var.storage_location

  force_destroy = false
  storage_class = "STANDARD"

  versioning {
    enabled = true
  }

  encryption {
    default_kms_key_name = google_kms_crypto_key.default.id
  }

  depends_on = [google_kms_crypto_key.default]
}

data "external" "public_files" {
  program     = ["python3", "scripts/list-files.py", "public"]
  working_dir = local.base_path
}

resource "google_storage_bucket_object" "public_files" {
  provider = google-beta

  for_each = data.external.public_files.result

  name   = each.value
  source = "${local.base_path}/${each.value}"

  bucket = local.firebase_bucket

  depends_on = [google_firebase_storage_bucket.default, data.external.public_files]
}

locals {
  public_dirs = distinct([for k, v in data.external.public_files.result : split("/", v)[1]])
}

output "frontend_public_files" {
  value = {
    for k, v in local.public_dirs : v => [for file in google_storage_bucket_object.public_files : file.name if split("/", file.name)[1] == v]
  }
}

resource "google_pubsub_topic" "donations_topic" {
  provider = google-beta
  project  = google_project.default.project_id
  name     = "donations"

  depends_on = [google_project_service.default]
}

resource "google_identity_platform_config" "default" {
  provider = google-beta
  project  = google_project.default.project_id

  autodelete_anonymous_users = true

  sign_in {
    allow_duplicate_emails = false

    anonymous {
      enabled = false
    }

    email {
      enabled           = true
      password_required = false
    }
  }

  blocking_functions {
    triggers {
      event_type   = "beforeCreate"
      function_uri = [for f in google_cloudfunctions2_function.cloud_functions : f.service_config[0].uri if f.name == "OnRegister"][0]
    }

    forward_inbound_credentials {
      id_token      = true
      access_token  = true
      refresh_token = true
    }
  }

  authorized_domains = [
    "localhost",
    "${google_project.default.project_id}.firebaseapp.com",
    "${google_project.default.project_id}.web.app"
  ]

  depends_on = [google_project_service.default, google_cloudfunctions2_function.cloud_functions]
}