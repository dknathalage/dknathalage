module "test" {
  source = "./modules/service-v1"

  name            = "test-${var.environment}"
  container_image = "gcr.io/cloudrun/hello"
  bucket = {
    location = "australia-southeast1"
  }
}
