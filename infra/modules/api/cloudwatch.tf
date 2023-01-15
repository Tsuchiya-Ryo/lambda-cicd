resource "aws_cloudwatch_log_group" "this" {
  name              = "/ecs/${var.api_name}"
  retention_in_days = "1"
}
