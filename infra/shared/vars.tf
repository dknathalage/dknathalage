variable "features" {
  default = {
    regions = [
      "us-central1",
      "europe-west1",
      "asia-southeast1",
      "asia-east1",
      "southamerica-east1",
      "australia-southeast1",
    ]

    dev_region = "australia-southeast1"

    registry_name = "dknathalage"
  }

  type = object({
    regions       = list(string)
    registry_name = string
  })
}
