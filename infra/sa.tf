resource "google_service_account" "gha_ci_sa" {
  account_id = "gha-ci-sa@dknathalage.iam.gserviceaccount.com"
}

resource "google_service_account_iam_binding" "workload_identity_user" {
  service_account_id = google_service_account.gha_ci_sa.name
  role               = "roles/iam.workloadIdentityUser"

  members = [
    "principalSet://iam.googleapis.com/projects/719430876063/locations/global/workloadIdentityPools/dknathalage-identity-pool/attribute.repository/dknathalage/dknathalage"
  ]
}

resource "google_service_account_iam_binding" "owner" {
  service_account_id = google_service_account.gha_ci_sa.name
  role               = "roles/owner"

  members = [
    "serviceAccount:gha-ci-sa@dknathalage.iam.gserviceaccount.com"
  ]
}