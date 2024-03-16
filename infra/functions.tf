resource "google_service_account" "account" {
  account_id   = "gcf-sa"
  display_name = "Cloud Functions Service Account"
  project      = google_project.default.project_id
  depends_on   = [google_project_service.default]
}


resource "google_project_iam_member" "event_receiver" {
  provider   = google-beta
  project    = google_project.default.project_id
  role       = "roles/eventarc.eventReceiver"
  member     = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_service_account.account]
}

resource "google_project_iam_member" "pubsub_publisher" {
  provider   = google-beta
  project    = google_project.default.project_id
  role       = "roles/pubsub.publisher"
  member     = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_service_account.account]
}

resource "google_project_iam_member" "kms_encrypter_decrypter" {
  provider   = google-beta
  project = google_project.default.project_id
  role    = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member  = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_service_account.account]
}

resource "google_project_iam_member" "firebase_admin" {
  provider   = google-beta
  project    = google_project.default.project_id
  role       = "roles/firebase.admin"
  member     = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_service_account.account]
}

resource "google_storage_bucket" "fn_bucket" {
  provider                    = google-beta
  project                     = google_project.default.project_id
  name                        = "${google_project.default.project_id}-cloud-source"
  location                    = var.storage_location
  uniform_bucket_level_access = true

  depends_on = [google_project_service.default]
}


data "archive_file" "source" {
  type        = "zip"
  source_dir  = "${local.base_path}/cloud"
  excludes    = ["cmd"]
  output_path = "${path.module}/.terraform/cloud_source.zip"
}

// The hash of the zip is needed to ensure that the functions are updated when the source changes
resource "google_storage_bucket_object" "functions_source" {
  name   = "cloud_source_${data.archive_file.source.output_sha256}.zip"
  source = data.archive_file.source.output_path
  bucket = google_storage_bucket.fn_bucket.name

  depends_on = [google_storage_bucket.fn_bucket, data.archive_file.source]
}

resource "google_cloudfunctions2_function" "user_register" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "UserRegister"
  location = var.gcf_location

  build_config {
    runtime     = "go121"
    entry_point = "UserRegister"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 0
    max_instance_count               = 10
    available_memory                 = "128Mi"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [google_storage_bucket_object.functions_source, google_service_account.account]
}

resource "google_cloud_run_service_iam_member" "user_register_invoker" {
  provider = google-beta
  project  = google_project.default.project_id

  service  = lower(google_cloudfunctions2_function.user_register.name)
  location = google_cloudfunctions2_function.user_register.location

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.user_register]
}

resource "google_cloudfunctions2_function" "user_edit" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "UserEdit"
  location = var.gcf_location

  build_config {
    runtime     = "go121"
    entry_point = "UserEdit"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 0
    max_instance_count               = 10
    available_memory                 = "128Mi"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [google_storage_bucket_object.functions_source, google_service_account.account]
}

resource "google_cloud_run_service_iam_member" "user_edit_invoker" {
  provider = google-beta
  project  = google_project.default.project_id

  service  = lower(google_cloudfunctions2_function.user_edit.name)
  location = google_cloudfunctions2_function.user_edit.location

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.user_edit]
}

resource "google_cloudfunctions2_function" "tip_send" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "TipSend"
  location = var.gcf_location

  build_config {
    runtime     = "go121"
    entry_point = "TipSend"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 0
    max_instance_count               = 10
    available_memory                 = "128Mi"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [google_storage_bucket_object.functions_source, google_service_account.account]
}

resource "google_cloud_run_service_iam_member" "tip_send_invoker" {
  provider = google-beta
  project  = google_project.default.project_id

  service  = lower(google_cloudfunctions2_function.tip_send.name)
  location = google_cloudfunctions2_function.tip_send.location

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.tip_send]
}

resource "google_cloudfunctions2_function" "tip_confirm" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "TipConfirm"
  location = var.gcf_location

  build_config {
    runtime     = "go121"
    entry_point = "TipConfirm"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 0
    max_instance_count               = 10
    available_memory                 = "128Mi"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [google_storage_bucket_object.functions_source, google_service_account.account, google_cloudfunctions2_function.user_register]
}

resource "google_cloud_run_service_iam_member" "tip_confirm_invoker" {
  provider = google-beta
  project  = google_project.default.project_id

  service  = lower(google_cloudfunctions2_function.tip_confirm.name)
  location = google_cloudfunctions2_function.tip_confirm.location

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.tip_confirm]
}

resource "google_cloudfunctions2_function" "twitch_webhook" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "TwitchWebhook"
  location = var.gcf_location

  build_config {
    runtime     = "go121"
    entry_point = "TwitchWebhook"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 1
    max_instance_count               = 10
    available_memory                 = "128Mi"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [
    google_storage_bucket_object.functions_source,
    google_service_account.account,
    google_pubsub_topic.events,
    google_cloudfunctions2_function.user_edit]
}

resource "google_cloud_run_service_iam_member" "twitch_webhook_invoker" {
  provider = google-beta
  project  = google_project.default.project_id

  service  = lower(google_cloudfunctions2_function.twitch_webhook.name)
  location = google_cloudfunctions2_function.twitch_webhook.location

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.twitch_webhook]
}

resource "google_cloudfunctions2_function" "on_event" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "OnEvent"
  location = var.gcf_location

  build_config {
    runtime     = "go121"
    entry_point = "OnEvent"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 1
    max_instance_count               = 10
    available_memory                 = "128Mi"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_INTERNAL_ONLY"
    service_account_email            = google_service_account.account.email
  }

  event_trigger {
    event_type            = "google.cloud.pubsub.topic.v1.messagePublished"
    pubsub_topic          = google_pubsub_topic.events.id
    retry_policy          = "RETRY_POLICY_DO_NOT_RETRY"
    service_account_email = google_service_account.account.email
  }

  depends_on = [google_storage_bucket_object.functions_source, google_service_account.account, google_pubsub_topic.events, google_cloudfunctions2_function.user_register]
}

resource "google_cloudfunctions2_function" "event_change_state" {
  provider = google-beta
  project  = google_project.default.project_id

  name     = "EventChangeState"
  location = var.gcf_location

  build_config {
    runtime     = "go121"
    entry_point = "EventChangeState"
    source {
      storage_source {
        bucket = google_storage_bucket.fn_bucket.name
        object = google_storage_bucket_object.functions_source.name
      }
    }
  }

  service_config {
    min_instance_count               = 0
    max_instance_count               = 10
    available_memory                 = "128Mi"
    max_instance_request_concurrency = 1
    timeout_seconds                  = 60
    ingress_settings                 = "ALLOW_ALL"
    service_account_email            = google_service_account.account.email
  }

  depends_on = [google_storage_bucket_object.functions_source, google_service_account.account, google_cloudfunctions2_function.user_edit]
}

resource "google_cloud_run_service_iam_member" "event_change_state_invoker" {
  provider = google-beta
  project  = google_project.default.project_id

  service  = lower(google_cloudfunctions2_function.event_change_state.name)
  location = google_cloudfunctions2_function.event_change_state.location

  member = "allUsers"
  role   = "roles/run.invoker"

  depends_on = [google_cloudfunctions2_function.event_change_state]
}