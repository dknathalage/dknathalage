resource "vultr_container_registry" "vcr1" {
  name   = "dknathalageregistry"
  region = "mel"
  plan   = "start_up"
  public = false
}
