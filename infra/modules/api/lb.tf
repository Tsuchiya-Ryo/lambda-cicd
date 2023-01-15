# ALB
resource "aws_lb" "main" {
  load_balancer_type = "application"
  name               = "${var.api_name}-alb"

  security_groups = ["${aws_security_group.alb.id}"]
  subnets         = ["${aws_subnet.public_1a.id}", "${aws_subnet.public_1c.id}"]
}


# Listener
resource "aws_lb_listener" "main" {
  port     = "80"
  protocol = "HTTP"

  load_balancer_arn = aws_lb.main.arn

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      status_code  = "200"
      message_body = "httpOk"
    }
  }
}


# ELB Target Group
resource "aws_lb_target_group" "main" {
  name = "${var.api_name}-tg"

  vpc_id = aws_vpc.main.id

  port        = 80
  protocol    = "HTTP"
  target_type = "ip"

  health_check {
    path                = "/health_check"
    healthy_threshold   = 5
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = 200
    port                = "traffic-port"
    protocol            = "HTTP"
  }
  depends_on = [
    aws_lb.main
  ]
}

# ALB Listener Rule
resource "aws_lb_listener_rule" "main" {
  listener_arn = aws_lb_listener.main.arn

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.arn
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}
