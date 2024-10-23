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
}

provider "google" {
  project = "dknathalage"
  region  = "australia-southeast2"

}

provider "vultr" {
  api_key     = "VULTR_API_KEY"
  rate_limit  = 100
  retry_limit = 3
}
