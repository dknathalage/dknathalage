resource "vultr_instance" "my_instance" {
  plan   = "vc2-1c-1gb"
  region = "mel"
  app_id = "1125"
}
