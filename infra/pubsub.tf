resource "google_pubsub_topic" "donations_topic" {
  provider = google-beta
  project = google_project.default.project_id
  name = "donations"

  depends_on = [google_project_service.default]
}
