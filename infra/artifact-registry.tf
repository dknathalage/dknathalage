variable "artifact_registries" {
  description = "A map of artifact registry configurations"
  default     = {}
  type = map(object({
    location      = string
    repository_id = string
  }))
}

resource "google_artifact_registry_repository" "docker-artifacts" {
  for_each      = var.artifact_registries
  location      = each.value.location
  repository_id = each.value.repository_id
  format        = "DOCKER"

  cleanup_policies {
    id     = "delete"
    action = "DELETE"
    condition {
      older_than            = "3600s"
      package_name_prefixes = []
      tag_prefixes          = []
      tag_state             = "ANY"
      version_name_prefixes = []
    }
  }

  cleanup_policies {
    id     = "keep"
    action = "KEEP"
    condition {
      tag_prefixes          = ["refs-"]
      tag_state             = "TAGGED"
      package_name_prefixes = []
      version_name_prefixes = []
    }
  }

  docker_config {
    immutable_tags = false
  }
}

import {
  for_each = var.artifact_registries
  id       = "dknathalage/${each.value.location}/${each.value.repository_id}"
  to       = google_artifact_registry_repository.docker-artifacts[each.key]
}
