variable "project_map" {
  description = "A map of project configurations"
  default     = {}
  type = map(object({
    project_id = string
  }))
}

resource "google_project" "my_project" {
  for_each   = var.project_map
  name       = each.key
  project_id = each.value.project_id
}

import {
  for_each = var.project_map
  id       = "dknathalage"
  to       = google_project.my_project[each.key]
}