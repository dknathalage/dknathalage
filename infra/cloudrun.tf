module "test" {
  source = "./modules/service-v1"

  name            = "test-${var.environment}"
  container_image = "us-docker.pkg.dev/cloudrun/container/hello"
  bucket = {
    location = "australia-southeast1"
  }

  project = var.project

  sa_roles = [
    "roles/run.invoker",
    "roles/storage.objectViewer",
    "storage.objects.list"
  ]
}
