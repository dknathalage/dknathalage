resource "vultr_instance" "test_instance" {
  plan     = "vc2-1c-1gb"
  region   = "mel"
  image_id = "docker"
  label    = "test-instance"
}
