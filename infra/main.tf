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
  base_path = "${path.module}/.."

  firebase_location   = "eur3"
  app_engine_location = "europe-west"
  storage_location    = "EU"
  functions_location  = "europe-west1"
  rtdb-location       = "europe-west1"
  kms_location        = "europe"

  go_runtime = "go121"
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
  ])

  service            = each.key
  disable_on_destroy = false
}

data "google_project" "project" {
  provider   = google-beta
  project_id = google_project.default.project_id
}

resource "google_project_iam_member" "default" {
  provider   = google-beta
  project    = data.google_project.project.project_id
  role       = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member     = "serviceAccount:service-${data.google_project.project.number}@gs-project-accounts.iam.gserviceaccount.com"
  depends_on = [google_project_service.default]
}

resource "google_kms_key_ring" "default" {
  provider = google-beta
  project  = google_project.default.project_id

  location   = local.kms_location
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
  location = local.storage_location

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