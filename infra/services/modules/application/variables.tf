variable "name" {
  type = string
}

variable "paths" {
  type = list(string)
}

variable "regions" {
  type = list(string)
}

variable "dockerfile" {
  type = string
}

variable "artifact_repository" {
  type = string
}

variable "project" {
  type = string
}

variable "service" {
  type = object({
    cpu    = number
    memory = number
    public = bool
  })
  nullable = true
}

variable "env" {
  type    = list(string)
  default = ["dev", "np", "prod"]
  validation {
    condition     = alltrue([for value in var.env : can(index(["dev", "np", "prod"], value) != -1)])
    error_message = "env must be either dev, np or prod"
  }
}
