resource "local_file" "ci_workflow" {
  content = yamlencode({
    name = "TF Gen ${var.name} Build 🏗️",
    on = {
      workflow_dispatch = {}
      workflow_call     = {}
    },
    env = {
      WORKLOAD_IDP = "projects/719430876063/locations/global/workloadIdentityPools/id-pool/providers/github-actions"
      CI_SA        = "gha-ci-sa@dknathalage.iam.gserviceaccount.com"
    }
    jobs = {
      build = {
        runs-on = "ubuntu-latest",
        steps = [
          {
            name = "Checkout 📥",
            uses = "actions/checkout@v4"
          },
          {
            name = "Set up Docker Buildx 🏗️",
            uses = "docker/setup-buildx-action@v3"
          },
          {
            name = "Login to Google Cloud 🌐",
            id   = "auth"
            uses = "google-github-actions/auth@v2"
            with = {
              token_format               = "access_token"
              workload_identity_provider = "$env.WORKLOAD_IDP"
              service_account            = "$env.CI_SA"
              access_token_lifetime      = "300s"
            }
          }
        ]
      }
    }
  })
  filename = "${path.cwd}/.github/workflows/tf_gen_${var.name}_ci.yml"
}

resource "local_file" "dev_cd_workflow" {
  count = can(index(var.env, "dev") != -1) ? 1 : 0
  content = yamlencode({
    name = "TF Gen ${var.name} Development Deployment 🚀",
    on = {
      workflow_dispatch = {}
      workflow_call     = {}
    },
    jobs = {
      deploy = {
        runs-on = "ubuntu-latest",
        steps = [
          {
            name = "Checkout",
            uses = "actions/checkout@v4"
          },
        ]
      }
    }
  })
  filename = "${path.cwd}/.github/workflows/tf_gen_${var.name}_dev_cd.yml"
}
