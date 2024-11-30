resource "google_iam_workload_identity_pool" "dkn_identity_pool" {
  workload_identity_pool_id = "id-pool"
}

import {
  id = "projects/dknathalage/locations/australia-southeast2/workloadIdentityPools/id-pool"
  to = google_iam_workload_identity_pool.dkn_identity_pool
}

resource "google_iam_workload_identity_pool_provider" "gha_provider" {
  project                            = "dknathalage"
  workload_identity_pool_id          = google_iam_workload_identity_pool.dkn_identity_pool.workload_identity_pool_id
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
  id = "projects/dknathalage/locations/australia-southeast2/workloadIdentityPools/id-pool/providers/github-actions"
  to = google_iam_workload_identity_pool_provider.gha_provider
}