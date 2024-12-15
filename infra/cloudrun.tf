module "test" {
  source = "./modules/service-v1"

  name = "invoissential-web-${var.environment}"

  artifact_registry = "dknathalage"
  image_name        = "invoissential/web:3984c182cf5c738fb9acee01994487bebd145adf"
  region            = "australia-southeast1"
  project           = var.project

  roles = [
    "roles/artifactregistry.reader",
    "roles/artifactregistry.writer",
  ]

  invokers = [
    "allUsers",
  ]
}
