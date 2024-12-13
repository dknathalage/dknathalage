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

variable "container_image" {
  type = string
}

resource "google_cloud_run_v2_service" "service" {
  name     = var.name
  location = var.location
  ingress  = var.ingress

  template {
    execution_environment = "EXECUTION_ENVIRONMENT_GEN2"

    containers {
      image = var.container_image

      dynamic "volume_mounts" {
        for_each = var.bucket != null ? [1] : []
        content {
          mount_path = "/mnt/data"
          name       = google_storage_bucket.bucket.name
        }
      }
    }

    dynamic "volumes" {
      for_each = var.bucket != null ? [1] : []
      content {
        name = google_storage_bucket.bucket.name
        gcs {
          bucket    = google_storage_bucket.bucket.name
          read_only = false
        }
      }
    }
  }
}
