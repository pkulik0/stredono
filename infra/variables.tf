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

variable "tenor_api_key" {
  description = "The api (public) key of the Tenor account"
  type        = string
  default     = "LIVDSRZULELA"
}

variable "fn_on_register" {
  description = "The function to call when a user registers"
  type        = string
  default     = "UserRegister"
}

variable "cloud_functions" {
  type = map(object({
    runtime       = string
    entry         = string
    public        = bool
    min_instances = number
    max_instances = number
    concurrency   = number
    memory        = string
    timeout       = number
    location      = string
  }))
  default = {
    UserRegister = {
      runtime       = "go121"
      entry         = "UserRegister"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    UserEdit = {
      runtime       = "go121"
      entry         = "UserEdit"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    TipSend = {
      runtime       = "go121"
      entry         = "TipSend"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    TipConfirm = {
      runtime       = "go121"
      entry         = "TipConfirm"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    TwitchCreateSub = {
      runtime       = "go121"
      entry         = "TwitchCreateSub"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    TwitchWebhook = {
      runtime       = "go121"
      entry         = "TwitchWebhook"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    TwitchGetData = {
      runtime       = "go121"
      entry         = "TwitchGetData"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
  }
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