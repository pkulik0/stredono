resource "google_service_account" "account" {
  account_id   = "gcf-sa"
  display_name = "Cloud Functions Service Account"
  project      = google_project.default.project_id
  depends_on   = [google_project_service.default]
}

resource "google_project_iam_member" "firebase_admin" {
  provider = google-beta
  project = google_project.default.project_id
  role    = "roles/firebase.admin"
  member  = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_service_account.account]
}

resource "google_storage_bucket" "fn_bucket" {
  provider                    = google-beta
  project                     = google_project.default.project_id
  name                        = "${google_project.default.project_id}-gcf-source"
  location                    = var.storage_location
  uniform_bucket_level_access = true

  depends_on = [google_project_service.default]
}

data "archive_file" "source" {
  type        = "zip"
  source_dir  = "${local.base_path}/cloud"
  excludes    = ["cmd"]
  output_path = "${path.module}/.terraform/functions_source.zip"
}

resource "google_storage_bucket_object" "functions_source" {
  name   = "functions_source.zip"
  source = data.archive_file.source.output_path
  bucket = google_storage_bucket.fn_bucket.name

  depends_on = [google_storage_bucket.fn_bucket, data.archive_file.source]
}

locals {
  public_functions = { for function in google_cloudfunctions2_function.cloud_functions : function.name => function.location if var.cloud_functions[function.name].public }
}

resource "google_cloudfunctions2_function" "cloud_functions" {
  for_each = var.cloud_functions

  provider = google-beta
  project  = google_project.default.project_id

  name     = each.key
  location = each.value.location

  build_config {
    runtime     = each.value.runtime
    entry_point = each.value.entry
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = each.value.min_instances
    max_instance_count               = each.value.max_instances
    available_memory                 = each.value.memory
    max_instance_request_concurrency = each.value.concurrency
    timeout_seconds                  = each.value.timeout
    ingress_settings                 = each.value.public ? "ALLOW_ALL" : "ALLOW_INTERNAL_ONLY"
    service_account_email            = google_service_account.account.email
  }

  lifecycle {

  }

  depends_on = [google_storage_bucket_object.functions_source, google_service_account.account]
}

// run.invoker for Gen 2 (used to be cloudfunctions.invoker for Gen 1)
resource "google_cloud_run_service_iam_member" "invoker" {
  for_each = local.public_functions

  provider = google-beta
  project  = google_project.default.project_id

  service  = lower(each.key)
  location = each.value

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.cloud_functions]
}