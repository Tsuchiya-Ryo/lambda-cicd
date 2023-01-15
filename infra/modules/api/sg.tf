# For ALB
resource "aws_security_group" "alb" {
  name        = "${var.api_name}-alb-sg"
  description = "security group for alb"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.api_name}-alb"
  }
}

# For ECS
resource "aws_security_group" "ecs" {
  name        = "${var.api_name}-ecs-sg"
  description = "security group for ecs"

  vpc_id = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"]
  }

  # ECRからのpullでも必要
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.api_name}-ecs"
  }
}
