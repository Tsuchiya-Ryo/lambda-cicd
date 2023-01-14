terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
  required_version = ">= 1.3.0"
}


module "repository" {
  source = "./modules/repository"
  aws_account_id = var.aws_account_id
  aws_region = var.aws_region
  function_name = var.function_name
  function_dir = var.function_dir
  s3_bucket_name = var.s3_bucket_name
}

module "lambda" {
  source = "./modules/lambda"
  function_name = var.function_name
  lambda_ecr_repository_url = module.repository.lambda_ecr_repository_url
  s3_bucket_arn = module.repository.s3_bucket_arn

  depends_on = [
    module.repository
  ]
}
