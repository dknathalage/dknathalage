resource "google_service_account" "service_account" {
  account_id = "gha-ci-sa@dknathalage.iam.gserviceaccount.com"
}

import {
  id = "gha-ci-sa@dknathalage.iam.gserviceaccount.com"
  to = google_service_account.service_account
}