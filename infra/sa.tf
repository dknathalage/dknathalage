
resource "google_service_account" "actions_sa" {
  count      = contains(var.environment, "prod") ? 1 : 0
  account_id = "gha-ci-sa"
}

resource "google_service_account_iam_member" "actions_sa_member" {
  count              = contains(var.environment, "prod") ? 1 : 0
  service_account_id = "projects/dknathalage/serviceAccounts/${google_service_account.actions_sa.email}"
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.dkn_identity_pool.name}/*"
}

resource "google_project_iam_binding" "actions_sa_binding" {
  count   = contains(var.environment, "prod") ? 1 : 0
  project = "dknathalage"
  role    = "roles/editor"
  members = [
    "serviceAccount:${google_service_account.actions_sa.email}"
  ]
}

import {

  to = google_service_account.actions_sa[0]
  id = "projects/dknathalage/serviceAccounts/gha-ci-sa@dknathalage.iam.gserviceaccount.com"
}

import {
  to = google_service_account_iam_member.actions_sa_member[0]
  id = "projects/dknathalage/serviceAccounts/gha-ci-sa@dknathalage.iam.gserviceaccount.com roles/iam.workloadIdentityUser principalSet://iam.googleapis.com/projects/719430876063/locations/global/workloadIdentityPools/id-pool/*"
}

import {
  to = google_project_iam_binding.actions_sa_binding[0]
  id = "dknathalage roles/editor serviceAccount:gha-ci-sa@dknathalage.iam.gserviceaccount.com"
}
