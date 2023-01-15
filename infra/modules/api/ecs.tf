# ECS Cluster
resource "aws_ecs_cluster" "main" {
  name = var.api_name
}

# ECS Service
resource "aws_ecs_service" "main" {
  name = var.api_name

  depends_on = [
    aws_lb_listener_rule.main
  ]

  cluster = aws_ecs_cluster.main.name

  launch_type = "FARGATE"

  desired_count = "1"

  task_definition = aws_ecs_task_definition.main.arn

  network_configuration {
    subnets          = ["${aws_subnet.public_1a.id}", "${aws_subnet.public_1c.id}"]
    security_groups  = ["${aws_security_group.ecs.id}"]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.main.arn
    container_name   = var.api_name
    container_port   = "80"
  }

}

# Task Definition
resource "aws_ecs_task_definition" "main" {
  family = var.api_name

  requires_compatibilities = ["FARGATE"]

  cpu    = "256"
  memory = "512"

  network_mode = "awsvpc"

  execution_role_arn = aws_iam_role.task_execution.arn
  task_role_arn      = aws_iam_role.task.arn

  container_definitions = jsonencode([
    {
      "name" : "${var.api_name}",
      "image" : "${var.api_ecr_repository_url}:latest",
      "essential" : true,
      "logConfiguration" : {
        "logDriver" : "awslogs",
        "options" : {
          "awslogs-region" : "ap-northeast-1",
          "awslogs-stream-prefix" : "api",
          "awslogs-group" : "/ecs/${var.api_name}"
        }
      },
      "portMappings" : [
        {
          "protocol" : "tcp",
          "containerPort" : 80
        }
      ]
    }
  ])
}
