output "ProjectId" {
  value = google_project.default.project_id
}

output "OnRegister_url" {
  value = google_cloudfunctions2_function.OnRegister.service_config[0].uri
}

output "SendDonate_url" {
  value = google_cloudfunctions2_function.SendDonate.service_config[0].uri
}

resource "local_file" "firebase_webapp_config" {
  filename = "${local.base_path}/app/src/lib/firebaseWebConfig.json"

  content  = jsonencode({
    apiKey             = data.google_firebase_web_app_config.default.api_key,
    authDomain         = data.google_firebase_web_app_config.default.auth_domain,
    databaseURL        = data.google_firebase_web_app_config.default.database_url,
    projectId          = google_project.default.project_id,
    storageBucket      = data.google_firebase_web_app_config.default.storage_bucket,
    messagingSenderId  = data.google_firebase_web_app_config.default.messaging_sender_id,
    appId              = data.google_firebase_web_app_config.default.web_app_id,
    measurementId      = data.google_firebase_web_app_config.default.measurement_id
  })

  depends_on = [data.google_firebase_web_app_config.default]
}

resource "local_file" "golang_constants" {
  filename = "${local.base_path}/functions/constants.go"

  content = <<EOF
package functions

const (
    ProjectID = "${google_project.default.project_id}"
    DatabaseURL = "${data.google_firebase_web_app_config.default.database_url}"
    GcSecretsPath = "projects/${google_project.default.number}/secrets"
)
EOF

  depends_on = [data.google_firebase_web_app_config.default, google_secret_manager_secret.secrets]
}
