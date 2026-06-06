# Your first Terraform config.
# The "local" provider lets you practice the workflow (init/plan/apply)
# without needing a cloud account or credentials.

terraform {
  required_version = ">= 1.0"

  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "~> 2.5"
    }
  }
}

# A variable you can override on the command line or in a .tfvars file.
variable "greeting" {
  type    = string
  default = "Terraform tracks my changes!"
}

# This resource creates a real file on disk when you run `terraform apply`.
resource "local_file" "hello" {
  content  = var.greeting
  filename = "${path.module}/hello.txt"
}

# Outputs print useful values after apply.
output "file_path" {
  value = local_file.hello.filename
}
