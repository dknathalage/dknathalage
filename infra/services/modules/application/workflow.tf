resource "local_file" "ci_workflow" {
  content = yamlencode({
    name = "TF Gen ${var.name} Build 🏗️",
    on = {
      workflow_dispatch = {}
      workflow_call     = {}
    },
    jobs = {
      build = {
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
