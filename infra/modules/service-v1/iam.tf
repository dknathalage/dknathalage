variable "sa_roles" {
  type    = list(string)
  default = []
}

resource "google_service_account" "service_account" {
  account_id   = "cloudrun-sa-${var.name}"
  display_name = "CloudRun Service Account for ${var.name}"
}
