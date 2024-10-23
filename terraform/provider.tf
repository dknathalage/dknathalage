terraform {
  required_providers {
    vultr = {
      source  = "vultr/vultr"
      version = "2.21.0"
    }

    google = {
      source  = "hashicorp/google"
      version = "6.8.0"
    }
  }

  backend "gcs" {
    bucket = "dknathalage-tfstate"
    prefix = "terraform/state"
  }
}

provider "google" {
  project = "dknathalage"
  region  = "australia-southeast2"
}

variable "VULTR_API_KEY" {
  type = string
}

provider "vultr" {
  api_key     = var.VULTR_API_KEY
  rate_limit  = 100
  retry_limit = 3
}
