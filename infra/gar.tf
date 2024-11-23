resource "google_artifact_registry_repository" "gar-dknathalage" {
  location      = "australia-southeast2"
  repository_id = "dknathalage"
  format        = "DOCKER"
}

import {
  id = "projects/dknathalage/locations/australia-southeast2/repositories/dknathalage"
  to = google_artifact_registry_repository.gar-dknathalage
}
