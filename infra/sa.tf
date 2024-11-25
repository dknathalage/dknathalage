resource "google_service_account" "service_account" {
  account_id = "gha-ci-sa"
}

resource "google_service_account_iam_binding" "service_account_iam_binding" {
  service_account_id = google_service_account.service_account.email
  role               = "roles/owner"

  members = [
    "serviceAccount: projects/dknathalage/serviceAccounts/gha-ci-sa@dknathalage.iam.gserviceaccount.com"
  ]
}

import {
  id = "projects/dknathalage/serviceAccounts/gha-ci-sa@dknathalage.iam.gserviceaccount.com"
  to = google_service_account.service_account
}

import {
  id = "projects/dknathalage/serviceAccounts/gha-ci-sa@dknathalage.iam.gserviceaccount.com roles/owner"
  to = google_service_account_iam_binding.service_account_iam_binding
}
