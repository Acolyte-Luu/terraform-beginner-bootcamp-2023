terraform {
    required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.18.1"
    }

  }
}

data "aws_caller_identity" "current" {
  
}