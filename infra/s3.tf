resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "slack-ai-lambda"
  acl    = "private"
}

resource "aws_s3_bucket_object" "lambda_code" {
  bucket       = aws_s3_bucket.lambda_bucket.id
  key          = "main.zip"
  source       = "../app/main.zip"
  acl          = "private"
  content_type = "application/zip"
}
