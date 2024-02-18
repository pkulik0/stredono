output "project_id" {
  value = google_project.default.project_id
}

output "on_register_url" {
  value = google_cloudfunctions2_function.OnRegister.service_config[0].uri
}

output "send_donate_url" {
  value = google_cloudfunctions2_function.SendDonate.service_config[0].uri
}

output "firebase_hosting_url" {
  value = google_firebase_hosting_site.default.default_url
}

output "firebase_webapp_config" {
  value      = data.google_firebase_web_app_config.default
  depends_on = [data.google_firebase_web_app_config.default]
}

output "recaptcha_site_key" {
  value      = element(split("/", google_recaptcha_enterprise_key.primary.id), length(split("/", google_recaptcha_enterprise_key.primary.id)) - 1)
  depends_on = [google_recaptcha_enterprise_key.primary]
}

output "project_number" {
  value      = google_project.default.number
  depends_on = [google_project.default]
}