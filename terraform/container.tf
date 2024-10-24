resource "vultr_container_registry" "vcr1" {
  name   = "dknathalage-registry"
  region = "mel"
  plan   = "start_up"
  public = false
}
