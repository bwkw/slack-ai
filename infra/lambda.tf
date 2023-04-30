resource "aws_lambda_function" "main" {
  function_name = "slack-ai"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.lambda_execution_role.arn

  s3_bucket = aws_s3_bucket.lambda_bucket.bucket
  s3_key    = aws_s3_object.lambda_code.key
}

resource "aws_iam_role" "lambda_execution_role" {
  name = "lambda_execution_role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "lambda_policy" {
  name   = "lambda_policy"
  role   = aws_iam_role.lambda_execution_role.id
  policy = aws_iam_policy.lambda_policy.policy
}

resource "aws_iam_policy" "lambda_policy" {
  name        = "lambda_policy"
  description = "IAM policy for Lambda execution"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:*:*:*"
      }
    ]
  })
}
