## What is

Question Reply Bot using API Gateway, Lambda, Go, SlackBot, OpenAI

## View

<img width="903" alt="image" src="https://github.com/bwkw/slack-ai/assets/63583536/2e4c1a50-0501-4315-b892-af9dc7e0e561">

## How to set up

1. Create Terraform secret variable file

```
cp .tfvars.example .tfvars
```

2. Set each value in .tfvars

```
SLACK_API_TOKEN=
SLACK_VERIFICATION_TOKEN=
OPENAI_API_KEY=
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
AWS_PROFILE={{your_profile}} terraform plan -var-file=.tfvars
```

5. Create resource

```
AWS_PROFILE={{your_profile}} terraform apply -var-file=.tfvars
```

6. Destroy resource

```
AWS_PROFILE={{your_profile}} terraform destroy -var-file=.tfvars
```

## Supplement

### Deploy s3 alone

```
AWS_PROFILE={{your_profile}} terraform plan -target=aws_s3_bucket.lambda_bucket -target=aws_s3_object.lambda_code
```
