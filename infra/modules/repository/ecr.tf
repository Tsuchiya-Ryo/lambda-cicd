resource "aws_ecr_repository" "lambda" {
  name                 = var.function_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "push_image" {
  depends_on = [
    aws_ecr_repository.lambda
  ]

  provisioner "local-exec" {
    command = "sh ${path.module}/build.sh"

    environment = {
      AWS_REGION      = var.aws_region
      AWS_ACCOUNT_ID  = var.aws_account_id
      FUNCTION_DIR  = var.function_dir
      REPOSITORY_NAME = aws_ecr_repository.lambda.name
      REPOSITORY_URL  = aws_ecr_repository.lambda.repository_url
    }
  }
}