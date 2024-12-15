module "test" {
  source = "./modules/service-v1"

  name              = "invoissential-web-${var.environment}"
  artifact_registry = "dknathalage"
  image_name        = "invoissential/web:${var.git_sha}"
  region            = "australia-southeast1"
  project           = var.project

  roles    = ["roles/artifactregistry.reader"]
  invokers = ["allUsers"]
}
