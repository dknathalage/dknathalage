
resource "google_service_account" "actions_sa" {
  account_id = "gha-ci-sa"
}

resource "google_service_account_iam_member" "actions_sa_member" {
  service_account_id = "projects/dknathalage/serviceAccounts/${google_service_account.actions_sa.email}"
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.dkn_identity_pool.name}/*"
}

resource "google_service_account_iam_member" "editor_permissions" {
  service_account_id = "projects/dknathalage/serviceAccounts/${google_service_account.actions_sa.email}"
  role               = "roles/editor"
  member             = "serviceAccount:${google_service_account.actions_sa.email}"
}
