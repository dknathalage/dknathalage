variable "sa_roles" {
  type    = list(string)
  default = []
}

resource "google_service_account" "service_account" {
  account_id   = "service-${var.name}-sa-${random_id.service_account_id.hex}"
  display_name = "CloudRun Service Account for ${var.name}"
}

resource "google_service_account_iam_member" "service-account-iam" {
  for_each           = { for role in var.sa_roles : role => role }
  service_account_id = google_service_account.service_account.email
  role               = each.value
  member             = "serviceAccount:${google_service_account.service_account.email}"
}

resource "random_id" "service_account_id" {
  byte_length = 8
}
