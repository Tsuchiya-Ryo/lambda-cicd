resource "aws_lambda_function" "test_lambda" {
  function_name = var.function_name
  architectures = ["x86_64"]
  package_type = "Image"
  image_uri = "${aws_ecr_repository.main.repository_url}:latest"
  role = aws_iam_role.iam_for_lambda.arn

  lifecycle {
    ignore_changes = [image_uri]
  }

  # 予めイメージがpushされていないとimage_uriの参照ができない
  depends_on = [
    null_resource.push_img
  ]
}
