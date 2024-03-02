locals {
  secrets = {
    "twitch-client-secret"   = {},
    "twitch-eventsub-secret" = {}
    "proxy"     = {}
  }
}

resource "google_secret_manager_secret" "secrets" {
  for_each = local.secrets

  provider  = google-beta
  project   = google_project.default.project_id
  secret_id = each.key

  replication {
    auto {}
  }

  depends_on = [google_project_service.default]
}

resource "google_secret_manager_secret_iam_member" "binding" {
  for_each = google_secret_manager_secret.secrets

  provider = google-beta
  project  = google_project.default.project_id

  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.account.email}"
  secret_id = each.value.secret_id

  depends_on = [google_secret_manager_secret.secrets]
}
