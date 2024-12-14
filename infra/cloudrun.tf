module "test" {
  source = "./modules/service-v1"

  name            = "invoissential-web-${var.environment}"
  container_image = "${google_artifact_registry_repository.docker-artifacts.id}/invoissential/web"
  bucket = {
    location = "australia-southeast1"
  }

  project = var.project

  sa_roles = [
    "roles/storage.objectViewer",
  ]

  invokers = [
    "allUsers",
  ]
}
