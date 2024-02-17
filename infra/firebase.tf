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

resource "google_firestore_index" "donations" {
  provider = google-beta
  project = google_project.default.project_id

  database = google_firestore_database.default.name
  collection = "donations"

  fields {
    field_path = "RecipientId"
    order = "ASCENDING"
  }

  fields {
    field_path = "Timestamp"
    order = "ASCENDING"
  }
}

resource "google_firestore_backup_schedule" "daily" {
  provider = google-beta
  project = google_project.default.project_id

  retention = "604800s"
  daily_recurrence {}

  depends_on = [google_firestore_database.default]
}

resource "google_firebaserules_ruleset" "firestore" {
  provider = google-beta
  project = google_project.default.project_id

  source {
    files {
      name = "firestore.rules"
      content = file("${local.base_path}/firestore.rules")
    }
  }

  depends_on = [google_firestore_database.default]
}

resource "google_firebaserules_release" "firestore" {
  provider = google-beta
  name = "cloud.firestore"
  ruleset_name = google_firebaserules_ruleset.firestore.name
  project = google_project.default.project_id

  depends_on = [google_firebaserules_ruleset.firestore]
}

resource "google_app_engine_application" "default-bucket" {
  provider = google-beta
  project = google_project.default.project_id

  location_id = local.app_engine_location

  depends_on = [google_firestore_database.default]
}

resource "google_firebase_storage_bucket" "default" {
  provider = google-beta
  project = google_project.default.project_id
  bucket_id = google_app_engine_application.default-bucket.default_bucket
}

locals {
  firebase_bucket = split("/", google_firebase_storage_bucket.default.name)[3]
}

resource "google_firebaserules_ruleset" "storage" {
  provider = google-beta
  project = google_project.default.project_id
  source {
    files {
      content = file("${local.base_path}/storage.rules")
      name    = "storage.rules"
    }
  }
  depends_on = [google_firebase_project.default]
}

resource "google_firebaserules_release" "storage" {
  provider = google-beta
  project = google_project.default.project_id
  name = "firebase.storage/${google_app_engine_application.default-bucket.default_bucket}"
  ruleset_name = "projects/${google_project.default.project_id}/rulesets/${google_firebaserules_ruleset.storage.name}"
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
    firebase_json_hash = filesha256("${local.base_path}/firebase.json")
    rtdb_rules_hash = filesha256("${local.base_path}/rtdb.rules.json")
  }

  depends_on = [
    google_firestore_database.default,
  ]

  provisioner "local-exec" {
    command = "firebase deploy --only database --project ${google_project.default.project_id}"
    working_dir = path.module
  }
}