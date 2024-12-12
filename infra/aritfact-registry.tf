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
}

import {
  for_each = var.artifact_registries
  to       = google_artifact_registry_repository.docker-artifacts
  id       = "${var.project}/${each.value.location}/${each.value.repository_id}"
}
