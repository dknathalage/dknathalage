variable "sa_roles" {
  type    = list(string)
  default = []
}

resource "google_service_account" "service_account" {
  account_id   = "${var.name}-sa"
  display_name = "CloudRun Service Account for ${var.name}"
}

resource "google_artifact_registry_repository_iam_member" "member" {
  project    = var.project
  location   = var.location
  repository = var.artifact_registry
  role       = "roles/artifactregistry.reader"
  member     = "serviceAccount:${google_service_account.service_account.email}"
}
