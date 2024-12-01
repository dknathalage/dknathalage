module "invoice" {
  source = "./modules/application"

  name                = "invoice"
  dockerfile          = "config/docker/invoice.Dockerfile"
  artifact_repository = "dknathalage"
  project             = "dknathalage"

  paths = [
    "config/**/invoice**",
    ".github/workflows/**",
    "modules/invoice/**",
  ]

  regions = [
    "australia-southeast-2",
  ]

  service = null
}
