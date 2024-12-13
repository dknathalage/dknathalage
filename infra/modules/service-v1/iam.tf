variable "sa_roles" {
  type    = list(string)
  default = []
}

resource "google_service_account" "service_account" {
  account_id   = "cloudrun-sa-${var.name}"
  display_name = "CloudRun Service Account for ${var.name}"
}

resource "google_service_account_iam_member" "service-account-iam" {
  for_each           = { for role in var.sa_roles : role => role }
  service_account_id = google_service_account.service_account.email
  role               = each.value
  member             = "serviceAccount:${google_service_account.service_account.email}"
}
