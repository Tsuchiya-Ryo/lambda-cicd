variable "aws_account_id" {}
variable "aws_region" {}
variable "function_name" {}
variable "docker_context" {}

resource "aws_ecr_repository" "main" {
    name = var.function_name
    image_tag_mutability = "MUTABLE"
  
    image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "push_img" {

  depends_on = [
    aws_ecr_repository.main
  ]
  
  # --password-stdinでのログインができなかった(windows wsl2)
  provisioner "local-exec" {
    command = "docker login --username AWS --password $(aws ecr get-login-password --region ${var.aws_region}) ${var.aws_account_id}.dkr.ecr.${var.aws_region}.amazonaws.com"
  }

  provisioner "local-exec" {
    command = "cd ${var.docker_context} && docker build -t ${aws_ecr_repository.main.name} ."
  }

  provisioner "local-exec" {
    command = "docker tag ${aws_ecr_repository.main.name}:latest ${aws_ecr_repository.main.repository_url}:latest"
  }

  provisioner "local-exec" {
    command = "docker push ${aws_ecr_repository.main.repository_url}:latest"
  }
}

