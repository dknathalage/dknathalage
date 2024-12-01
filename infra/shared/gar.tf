resource "google_artifact_registry_repository" "gar" {
  for_each               = toset(var.features.regions)
  location               = each.key
  repository_id          = var.features.registry_name
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
