
resource "google_service_account" "actions_sa" {
  count      = contains(["prod"], var.environment) ? 1 : 0
  account_id = "gha-ci-sa"
}

resource "google_service_account_iam_member" "actions_sa_member" {
  count              = contains(["prod"], var.environment) ? 1 : 0
  service_account_id = "projects/dknathalage/serviceAccounts/${google_service_account.actions_sa[0].email}"
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.dkn_identity_pool[0].name}/*"
}

resource "google_project_iam_binding" "actions_sa_binding" {
  count   = contains(["prod"], var.environment) ? 1 : 0
  project = "dknathalage"
  role    = "roles/editor"
  members = [
    "serviceAccount:${google_service_account.actions_sa[0].email}"
  ]
}

import {
  count = contains(["prod"], var.environment) ? 1 : 0
  to    = google_service_account.actions_sa[0]
  id    = "projects/dknathalage/serviceAccounts/gha-ci-sa@dknathalage.iam.gserviceaccount.com"
}

import {
  count = contains(["prod"], var.environment) ? 1 : 0
  to    = google_service_account_iam_member.actions_sa_member[0]
  id    = "projects/dknathalage/serviceAccounts/gha-ci-sa@dknathalage.iam.gserviceaccount.com roles/iam.workloadIdentityUser principalSet://iam.googleapis.com/projects/719430876063/locations/global/workloadIdentityPools/id-pool/*"
}

import {
  count = contains(["prod"], var.environment) ? 1 : 0
  to    = google_project_iam_binding.actions_sa_binding[0]
  id    = "dknathalage roles/editor serviceAccount:gha-ci-sa@dknathalage.iam.gserviceaccount.com"
}

resource "google_iam_workload_identity_pool" "dkn_identity_pool" {
  count                     = contains(["prod"], var.environment) ? 1 : 0
  workload_identity_pool_id = "id-pool"
}

import {
  count = contains(["prod"], var.environment) ? 1 : 0
  id    = "projects/dknathalage/locations/global/workloadIdentityPools/id-pool"
  to    = google_iam_workload_identity_pool.dkn_identity_pool[0]
}

resource "google_iam_workload_identity_pool_provider" "gha_provider" {
  count                              = contains(["prod"], var.environment) ? 1 : 0
  project                            = "dknathalage"
  workload_identity_pool_id          = google_iam_workload_identity_pool.dkn_identity_pool[0].workload_identity_pool_id
  workload_identity_pool_provider_id = "github-actions"
  attribute_condition                = "attribute.repository_owner==\"dknathalage\""
  attribute_mapping = {
    "google.subject"             = "assertion.sub"
    "attribute.actor"            = "assertion.actor"
    "attribute.aud"              = "assertion.aud"
    "attribute.repository"       = "assertion.repository"
    "attribute.repository_owner" = "assertion.repository_owner"
  }
  oidc {
    allowed_audiences = []
    issuer_uri        = "https://token.actions.githubusercontent.com"
  }
}

import {
  count = contains(["prod"], var.environment) ? 1 : 0
  id    = "projects/dknathalage/locations/global/workloadIdentityPools/id-pool/providers/github-actions"
  to    = google_iam_workload_identity_pool_provider.gha_provider[0]
}
