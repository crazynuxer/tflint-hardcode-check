resource "aws_s3_bucket" "example" {
  bucket = "example-bucket"

  lifecycle_rule {
    id      = "rule-1"
    enabled = true

    expiration {
      days = 30
    }
  }

  tags = {
    Name = "example-bucket"
    ARN  = "arn:aws:s3:::example-bucket" # Hardcoded ARN
  }
}

