resource "google_pubsub_schema" "events" {
  provider = google-beta
  project  = google_project.default.project_id

  name       = "events"
  definition = data.local_file.events_pb.content
  type       = "PROTOCOL_BUFFER"

  depends_on = [google_project_service.default, data.local_file.events_pb]
}

resource "google_pubsub_topic" "events" {
  provider = google-beta
  project  = google_project.default.project_id
  name     = "events"

  schema_settings {
    schema   = google_pubsub_schema.events.id
    encoding = "BINARY"
  }

  depends_on = [google_project_service.default, google_pubsub_schema.events]
}