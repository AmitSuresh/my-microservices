
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# CloudWatch Log Group Resource
resource "aws_cloudwatch_log_group" "order_log_group" {
  name              = "/aws/order-service/logs"
  retention_in_days = 30
}
