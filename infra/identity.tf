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
      function_uri = google_cloudfunctions2_function.OnRegister.service_config[0].uri
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

  depends_on = [google_project_service.default, google_cloudfunctions2_function.OnRegister]
}