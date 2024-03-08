variable "is_local" {
  description = "Whether the deployment is done locally or not"
  type        = bool
  default     = true
}

locals {
  domains = [
    "${google_project.default.project_id}.firebaseapp.com",
    "${google_project.default.project_id}.web.app",
    "stredono.com",
  ]
}

variable "twitch_client_id" {
  description = "The client id of the Twitch application"
  type        = string
  default     = "t1kl0vkt6hv06bi4ah4691hi8fexso"
}

variable "twitch_uid" {
  description = "The user id of the brand Twitch account"
  type        = string
  default     = "1033918710"
}

variable "tenor_api_key" {
  description = "The api (public) key of the Tenor account"
  type        = string
  default     = "LIVDSRZULELA"
}

variable "gcf_location" {
  description = "The default region of gcp cloud functions"
  type        = string
  default     = "europe-west1"
}

variable "firebase_location" {
  description = "The location of the Firebase project"
  type        = string
  default     = "eur3"
}

variable "app_engine_location" {
  description = "The location of the App Engine project"
  type        = string
  default     = "europe-west"
}

variable "storage_location" {
  description = "The location of the storage"
  type        = string
  default     = "EU"
}

variable "rtdb_location" {
  description = "The location of the real-time database"
  type        = string
  default     = "europe-west1"
}

variable "kms_location" {
  description = "The location of the key management service"
  type        = string
  default     = "europe"
}