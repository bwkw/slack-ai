## Overview

Question Reply Bot using API Gateway, Lambda, Go, SlackBot, OpenAI

## How to set up

1. Create Terraform secret variable file

```
cp .tfvars.example .tfvars
```

2. Set each value in .tfvars

```
SLACK_API_TOKEN = {{SlackのAPIトークン}}
SLACK_VERIFICATION_TOKEN = {{SlackAPIの認証用トークン}}
OPENAI_API_KEY = {{OpenAIのAPIキー}}
```

3. Obtain AWS access_key and secret_key

4. Set up the profile in the AWS CLI

```
aws configure --profile <profile_name>
```

## Deploy

1. Init workspace

```
terraform init
```

2. Format tf file

```
terraform fmt
```

3. Validate tf file

```
terraform validate
```

4. Check what will be created

```
AWS_PROFILE={{your_profile}} terraform plan
```

5. Create resource

```
AWS_PROFILE={{your_profile}} terraform apply
```

6. Destroy resource

```
AWS_PROFILE={{your_profile}} terraform destroy
```

## Supplement

### Deploy s3 alone.

```
AWS_PROFILE={{your_profile}} terraform plan -target=aws_s3_bucket.lambda_bucket -target=aws_s3_object.lambda_code
```
