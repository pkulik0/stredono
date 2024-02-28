output "backend_project_id" {
  value      = google_project.default.project_id
  depends_on = [google_project.default]
}

output "backend_project_number" {
  value      = google_project.default.number
  depends_on = [google_project.default]
}

output "backend_firebase_storage_bucket" {
  value = data.google_firebase_web_app_config.default.storage_bucket
}

output "frontend_function_urls" {
  value = { for function in google_cloudfunctions2_function.cloud_functions : function.name => function.service_config[0].uri if var.cloud_functions[function.name].public }
}

output "firebase_hosting_url" {
  value      = google_firebase_hosting_site.default.default_url
  depends_on = [google_firebase_hosting_site.default]
}

output "frontend_firebase_webapp_config" {
  value = {
    apiKey            = data.google_firebase_web_app_config.default.api_key,
    authDomain        = data.google_firebase_web_app_config.default.auth_domain,
    databaseURL       = data.google_firebase_web_app_config.default.database_url,
    projectId         = google_project.default.project_id,
    storageBucket     = data.google_firebase_web_app_config.default.storage_bucket,
    messagingSenderId = data.google_firebase_web_app_config.default.messaging_sender_id,
    appId             = data.google_firebase_web_app_config.default.web_app_id,
    measurementId     = data.google_firebase_web_app_config.default.measurement_id
  }
  description = "Firebase web app configuration in JSON format."
}

output "backend_firebase_database_url" {
  value = data.google_firebase_web_app_config.default.database_url
}

output "frontend_recaptcha_site_key" {
  value      = element(split("/", google_recaptcha_enterprise_key.primary.id), length(split("/", google_recaptcha_enterprise_key.primary.id)) - 1)
  depends_on = [google_recaptcha_enterprise_key.primary]
}

output "frontend_tenor_api_key" {
  value = var.tenor_api_key
}