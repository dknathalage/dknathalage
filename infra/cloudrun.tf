module "test" {
  source = "./modules/service-v1"

  name            = "invoissential-web-${var.environment}"
  container_image = "australia-southeast1-docker.pkg.dev/dknathalage/dknathalage/invoissential/web"
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
