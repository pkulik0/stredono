terraform {
  required_providers {
    google-beta = {
      source = "hashicorp/google-beta"
      version = "~> 5.0"
    }
  }
}

locals {
  firebase_location = "eur3"
  app_engine_location = "europe-west"
  storage_location = "EU"
  functions_location = "europe-west1"
  rtdb-location = "europe-west1"

  go_runtime = "go121"
}

provider "google-beta" {
  user_project_override = true
}

provider "google-beta" {
  alias = "no_user_project_override"
  user_project_override = false
}

resource "random_id" "project-id" {
  byte_length = 4
}

resource "google_project" "default" {
  provider = google-beta.no_user_project_override
  name       = "Stredono"
  project_id = "stredono-${random_id.project-id.hex}"

  billing_account = "0152B3-F5E957-35889E"

  labels = {
    "firebase" = "enabled"
  }
}

resource "google_project_service" "default" {
  provider = google-beta.no_user_project_override
  project = google_project.default.project_id
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
  ])
  service = each.key
  disable_on_destroy = false
}

resource "google_firebase_project" "default" {
  provider = google-beta
  project = google_project.default.project_id

  depends_on = [
    google_project_service.default
  ]
}

resource "google_firestore_database" "default" {
  provider = google-beta
  project = google_project.default.project_id
  name = "(default)"
  location_id = local.firebase_location
  type ="FIRESTORE_NATIVE"
  concurrency_mode = "OPTIMISTIC"

  depends_on = [google_firebase_project.default]
}

resource "google_firebaserules_ruleset" "default" {
  provider = google-beta
  project = google_project.default.project_id

  source {
    files {
      name = "firestore.rules"
      content = file("firestore.rules")
    }
  }

  depends_on = [google_firestore_database.default]
}

resource "google_firebaserules_release" "default" {
  provider = google-beta
  name = "cloud.firestore"
  ruleset_name = google_firebaserules_ruleset.default.name
  project = google_project.default.project_id

  depends_on = [google_firestore_database.default]
}

resource "google_app_engine_application" "default-bucket-fs" {
  provider = google-beta
  project = google_project.default.project_id

  location_id = local.app_engine_location

  depends_on = [google_firestore_database.default]
}

resource "google_firebase_storage_bucket" "default-bucket-fs" {
  provider = google-beta
  project = google_project.default.project_id
  bucket_id = google_app_engine_application.default-bucket-fs.default_bucket
}

resource "google_firebaserules_ruleset" "default-bucket-fs" {
  provider = google-beta
  project = google_project.default.project_id
  source {
    files {
      content = file("storage.rules")
      name    = "storage.rules"
    }
  }
  depends_on = [google_firebase_project.default]
}

resource "google_firebaserules_release" "default-bucket-fs" {
  provider = google-beta
  project = google_project.default.project_id
  name = "firebase.storage/${google_app_engine_application.default-bucket-fs.default_bucket}"
  ruleset_name = "projects/${google_project.default.project_id}/rulesets/${google_firebaserules_ruleset.default-bucket-fs.name}"
}

resource "google_firebase_database_instance" "default" {
  provider = google-beta
  project = google_project.default.project_id

  region = local.rtdb-location
  instance_id = "${google_project.default.project_id}-default-rtdb"
  type = "DEFAULT_DATABASE"

  depends_on = [google_firebase_project.default]
}

// Realtime Database Rules aren't supported by the google-beta provider yet, so we use the firebase CLI to deploy them
resource "null_resource" "run_firebase_deploy" {
  triggers = {
    firebase_json_hash = filesha256("${path.module}/firebase.json")
    rtdb_rules_hash = filesha256("${path.module}/rtdb.rules.json")
  }

  depends_on = [
    google_firestore_database.default,
  ]

  provisioner "local-exec" {
    command = "firebase deploy --only database --project ${google_project.default.project_id}"
    working_dir = path.module
  }
}

resource "google_service_account" "account" {
  account_id = "gcf-sa"
  display_name = "Cloud Functions Service Account"
  project = google_project.default.project_id
}

resource "google_pubsub_topic" "donations_topic" {
  provider = google-beta
  project = google_project.default.project_id
  name = "donations"

  depends_on = [google_project_service.default]
}

resource "google_storage_bucket" "fn_bucket" {
  provider = google-beta
  project = google_project.default.project_id
  name = "${google_project.default.project_id}-gcf-source"
  location = local.storage_location
  uniform_bucket_level_access = true

  depends_on = [google_project_service.default]
}

data "archive_file" "source" {
  type = "zip"
  source_dir = "${path.module}/functions"
  output_path = "${path.module}/functions_source.zip"
}

resource "google_storage_bucket_object" "functions_source" {
  name = "functions_source.zip"
  source = data.archive_file.source.output_path
  bucket = google_storage_bucket.fn_bucket.name

  depends_on = [google_storage_bucket.fn_bucket, data.archive_file.source]
}

resource "google_cloudfunctions2_function" "OnRegister" {
  provider = google-beta
  project = google_project.default.project_id

  name = "OnRegister"
  location = local.functions_location

  build_config {
    runtime = local.go_runtime
    entry_point = "OnRegister"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count = 0
    max_instance_count = 1
    available_memory = "256M"
    max_instance_request_concurrency = 1
    timeout_seconds = 60
    ingress_settings = "ALLOW_INTERNAL_ONLY"
    service_account_email = google_service_account.account.email
  }

  depends_on = [
    google_project_service.default
  ]
}

output "OnRegister_url" {
  value = google_cloudfunctions2_function.OnRegister.service_config[0].uri
}

resource "google_cloudfunctions2_function" "SendDonate" {
  provider = google-beta
  project = google_project.default.project_id

  name = "SendDonate"
  location = local.functions_location

  build_config {
    runtime = local.go_runtime
    entry_point = "SendDonate"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count = 0
    max_instance_count = 1
    available_memory = "256M"
    max_instance_request_concurrency = 1
    timeout_seconds = 60
    ingress_settings = "ALLOW_ALL"
    service_account_email = google_service_account.account.email
  }

  depends_on = [google_project_service.default]
}

output "SendDonate_url" {
  value = google_cloudfunctions2_function.SendDonate.service_config[0].uri
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"

  depends_on = [google_project_service.default]
}

resource "google_identity_platform_config" "default" {
  provider = google-beta
  project = google_project.default.project_id

  autodelete_anonymous_users = true

  sign_in {
    allow_duplicate_emails = false

    anonymous {
      enabled = false
    }

    email {
      enabled = true
      password_required = false
    }
  }

  blocking_functions {
    triggers {
      event_type   = "beforeCreate"
      function_uri = google_cloudfunctions2_function.OnRegister.service_config[0].uri
    }

    forward_inbound_credentials {
      id_token = true
      access_token = true
      refresh_token = true
    }
  }

  authorized_domains = [
    "localhost",
    "${google_project.default.project_id}.firebaseapp.com",
    "${google_project.default.project_id}.web.app"
  ]

  depends_on = [google_project_service.identitytoolkit, google_cloudfunctions2_function.OnRegister]
}

resource "google_firebase_web_app" "stredono_web" {
  provider = google-beta
  project = google_project.default.project_id

  display_name = "Stredono Web"

  deletion_policy = "DELETE"
  depends_on = [google_firebase_project.default]
}