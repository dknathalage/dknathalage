variable "bucket" {
  type = object({
    location = string
  })
  default = null
}

resource "google_storage_bucket" "bucket" {
  count = var.bucket != null ? 1 : 0

  name     = "service-${var.name}-bucket-${random_id.bucket_id[0].hex}"
  location = var.bucket.location
}

resource "random_id" "bucket_id" {
  count       = var.bucket != null ? 1 : 0
  byte_length = 8
}

resource "google_storage_bucket_iam_member" "bucket_reader" {
  count  = var.bucket != null ? 1 : 0
  bucket = google_storage_bucket.bucket[0].name
  role   = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.service_account.email}"
}
