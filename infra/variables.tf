variable "is_local" {
  description = "Whether the deployment is local or not"
  type        = bool
  default     = true
}

variable "tenor_api_key" {
  description = "The api (public) key of the Tenor account"
  type        = string
  default     = "LIVDSRZULELA"
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
    OnRegister = {
      runtime       = "go121"
      entry         = "OnRegister"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    SendTip = {
      runtime       = "go121"
      entry         = "SendTip"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    ConfirmPayment = {
      runtime       = "go121"
      entry         = "ConfirmPayment"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    GetListeners = {
      runtime       = "go121"
      entry         = "GetListeners"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    TwitchConnect = {
      runtime       = "go121"
      entry         = "TwitchConnect"
      public        = true
      min_instances = 0
      max_instances = 1
      concurrency   = 1
      memory        = "256M"
      timeout       = 60
      location      = "europe-west1"
    }
    TwitchSubscribe = {
      runtime       = "go121"
      entry         = "TwitchSubscribe"
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