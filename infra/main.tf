terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
  required_version = ">= 1.3.0"
}

module "lambda" {
  source         = "./modules/lambda"
  aws_account_id = var.aws_account_id
  aws_region     = var.aws_region
  function_name  = var.function_name
  docker_context = var.docker_context
}
