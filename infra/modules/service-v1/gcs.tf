variable "bucket" {
  type = object({
    location = string
  })
  default = null
}

resource "google_storage_bucket" "bucket" {
  count = var.bucket != null ? 1 : 0

  name     = "service-${var.name}-bucket-${random_id.bucket_id.hex}"
  location = var.bucket.location
}

resource "random_id" "bucket_id" {
  count       = var.bucket != null ? 1 : 0
  byte_length = 8
}
