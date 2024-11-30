resource "google_artifact_registry_repository" "gar-dknathalage" {
  for_each               = toset(["australia-southeast2", "us-central1"])
  location               = each.key
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
    immutable_tags = false
  }
}
