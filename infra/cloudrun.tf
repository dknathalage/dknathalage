module "test" {
  source = "./modules/service-v1"

  name            = "test"
  container_image = "australia-southeast1-docker.pkg.dev/dknathalage/dknathalage/invoissential/web"
  bucket = {
    location = "australia-southeast1"
  }
}
