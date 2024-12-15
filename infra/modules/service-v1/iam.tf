variable "sa_roles" {
  type    = list(string)
  default = []
}

resource "google_service_account" "service_account" {
  account_id   = "${var.name}-sa"
  display_name = "CloudRun Service Account for ${var.name}"
}

output "service_account" {
  value = google_service_account.service_account
}
