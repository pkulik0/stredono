terraform {
  required_providers {
    google-beta = {
      source = "hashicorp/google-beta"
      version = "~> 5.0"
    }
  }
}

locals {
  base_path = "${path.module}/.."

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
    "firebaseappcheck.googleapis.com",
    "secretmanager.googleapis.com",
    "recaptchaenterprise.googleapis.com",
  ])

  service = each.key
  disable_on_destroy = false
}