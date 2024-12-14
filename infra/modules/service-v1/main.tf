variable "name" {
  type = string
}

variable "location" {
  type    = string
  default = "australia-southeast1"
}

variable "ingress" {
  type    = string
  default = "INGRESS_TRAFFIC_ALL"
}

variable "artifact_registry" {
  type = string
}

variable "image_name" {
  type = string
}

variable "region" {
  type = string
}

variable "project" {
  type = string
}

variable "invokers" {
  type    = list(string)
  default = []
}

resource "google_cloud_run_v2_service" "service" {
  name                = var.name
  location            = var.location
  project             = var.project
  ingress             = var.ingress
  deletion_protection = false

  template {
    execution_environment = "EXECUTION_ENVIRONMENT_GEN2"
    service_account       = google_service_account.service_account.email

    containers {
      image = "${var.region}-docker.pkg.dev/${var.project}/${var.artifact_registry}/${var.image_name}"

      volume_mounts {
        mount_path = "/mnt/data"
        name       = google_storage_bucket.bucket.name
      }
    }

    volumes {
      name = google_storage_bucket.bucket.name
      gcs {
        bucket    = google_storage_bucket.bucket.name
        read_only = false
      }
    }
  }
}

resource "google_cloud_run_v2_service_iam_member" "invokers" {
  for_each = { for i, value in var.invokers : i => value }
  name     = google_cloud_run_v2_service.service.name
  location = google_cloud_run_v2_service.service.location
  project  = google_cloud_run_v2_service.service.project
  role     = "roles/run.invoker"
  member   = each.value
}
