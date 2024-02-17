resource "google_service_account" "account" {
  account_id   = "gcf-sa"
  display_name = "Cloud Functions Service Account"
  project      = google_project.default.project_id
}

resource "google_storage_bucket" "fn_bucket" {
  provider                    = google-beta
  project                     = google_project.default.project_id
  name                        = "${google_project.default.project_id}-gcf-source"
  location                    = local.storage_location
  uniform_bucket_level_access = true

  depends_on = [google_project_service.default]
}

data "archive_file" "source" {
  type        = "zip"
  source_dir  = "${local.base_path}/functions"
  excludes    = ["cmd"]
  output_path = "${path.module}/functions_source.zip"
}

resource "google_storage_bucket_object" "functions_source" {
  name   = "functions_source.zip"
  source = data.archive_file.source.output_path
  bucket = google_storage_bucket.fn_bucket.name

  depends_on = [google_storage_bucket.fn_bucket, data.archive_file.source]
}

resource "google_cloudfunctions2_function" "OnRegister" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "onregister"
  location = local.functions_location

  build_config {
    runtime     = local.go_runtime
    entry_point = "OnRegister"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 0
    max_instance_count               = 1
    available_memory                 = "256M"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [google_project_service.default, google_storage_bucket.fn_bucket, data.archive_file.source]
}

// Cloud Run (v1) Invoker role for Gen 2 (used to be cloudfunctions.invoker for Gen 1)
resource "google_cloud_run_service_iam_member" "invoker" {
  provider = google-beta
  project  = google_project.default.project_id

  location = google_cloudfunctions2_function.OnRegister.location
  service  = google_cloudfunctions2_function.OnRegister.name

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.OnRegister]
}

resource "google_cloudfunctions2_function" "SendDonate" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "SendDonate"
  location = local.functions_location

  build_config {
    runtime     = local.go_runtime
    entry_point = "SendDonate"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 0
    max_instance_count               = 1
    available_memory                 = "256M"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [google_project_service.default]
}