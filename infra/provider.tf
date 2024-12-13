provider "google" {
  project = "dknathalage"
  region  = "australia-southeast2"
}

terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.12.0"
    }
  }
}

terraform {
  backend "gcs" {
    bucket = "dknathalage-tfstate"
  }
}
