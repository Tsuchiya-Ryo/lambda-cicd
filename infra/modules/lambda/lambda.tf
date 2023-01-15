resource "aws_lambda_function" "main" {
  function_name = var.function_name
  architectures = ["x86_64"]
  package_type  = "Image"
  image_uri     = "${var.lambda_ecr_repository_url}:latest"
  role          = aws_iam_role.lambda.arn
  memory_size   = 1024
  timeout       = 10

  lifecycle {
    ignore_changes = [image_uri]
  }

  depends_on = [
    aws_cloudwatch_log_group.this
  ]
}
