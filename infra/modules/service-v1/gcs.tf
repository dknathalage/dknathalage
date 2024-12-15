resource "google_storage_bucket" "bucket" {
  name     = "service-${var.name}-bucket-${random_id.bucket_id.hex}"
  location = var.location
}

resource "random_id" "bucket_id" {
  byte_length = 8
}

resource "google_storage_bucket_iam_member" "bucket_reader" {
  bucket = google_storage_bucket.bucket.name
  role   = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.service_account.email}"
}
