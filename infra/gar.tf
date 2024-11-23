resource "google_artifact_registry_repository" "gar-dknathalage" {
  location               = "australia-southeast2"
  repository_id          = "dknathalage"
  format                 = "DOCKER"
  cleanup_policy_dry_run = false

  cleanup_policies {
    id     = "delete"
    action = "DELETE"
    condition {
      older_than = "3600s"
      tag_state  = "ANY"
    }
  }

  cleanup_policies {
    action = "KEEP"
    id     = "keep"
    condition {
      tag_prefixes = [
        "refs-",
      ]
      tag_state = "TAGGED"
    }
  }

  docker_config {
    immutable_tags = true
  }
}

import {
  id = "projects/dknathalage/locations/australia-southeast2/repositories/dknathalage"
  to = google_artifact_registry_repository.gar-dknathalage
}
