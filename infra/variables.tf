variable "environment" {
  description = "The environment to deploy the infrastructure to"
  type        = string
}

variable "project" {
  description = "The project to deploy the infrastructure to"
  type        = string
}

variable "git_sha" {
  description = "The git sha of trigger"
  default     = ""
  type        = string
}
