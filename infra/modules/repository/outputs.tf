output "lambda_ecr_repository_url" {
  value = aws_ecr_repository.lambda.repository_url
}

output "api_ecr_repository_url" {
  value = aws_ecr_repository.api.repository_url
}

output "s3_bucket_arn" {
  value = aws_s3_bucket.main.arn
}