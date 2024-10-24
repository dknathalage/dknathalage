resource "vultr_container_registry" "vcr1" {
  name   = "dknathalageregistry"
  region = "sjc"
  plan   = "start_up"
  public = false
}
