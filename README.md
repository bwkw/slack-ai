## Overview

Question Reply Bot using API Gateway, Lambda, Go, SlackBot, OpenAI

## How to set up

### Deploy s3 alone.

```
AWS_PROFILE={{your_profile}} terraform plan -target=aws_s3_bucket.lambda_bucket -target=aws_s3_bucket_object.lambda_code
```
