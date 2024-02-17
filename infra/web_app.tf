resource "google_firebase_web_app" "stredono_web" {
  provider = google-beta
  project  = google_project.default.project_id

  display_name = "Stredono Web"

  deletion_policy = "DELETE"
  depends_on      = [google_firebase_project.default]
}

data "google_firebase_web_app_config" "default" {
  provider = google-beta
  project  = google_project.default.project_id

  web_app_id = google_firebase_web_app.stredono_web.app_id

  depends_on = [google_firebase_web_app.stredono_web]
}

resource "google_recaptcha_enterprise_key" "primary" {
  provider     = google-beta
  project      = google_project.default.project_id
  display_name = "Stredono Main Key"

  web_settings {
    integration_type  = "SCORE"
    allowed_domains   = []
    allow_all_domains = true
  }

  depends_on = [google_project_service.default]
}

resource "google_firebase_app_check_service_config" "default" {
  provider = google-beta
  project  = google_project.default.project_id

  for_each = toset([
    "firebasestorage.googleapis.com",
    "firebasedatabase.googleapis.com",
    "firestore.googleapis.com",
    "identitytoolkit.googleapis.com",
  ])

  service_id       = each.key
  enforcement_mode = "ENFORCED"

  depends_on = [google_recaptcha_enterprise_key.primary]
}