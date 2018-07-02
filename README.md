# aws-sam-with-sfn

## Requirements

* [AWS SAM](https://github.com/awslabs/serverless-application-model)

## Super Simple Deployment

```bash
make build

sam package \
    --template-file ./template.yaml \
    --output-template-file pkg/packaged-template.yaml \
    --s3-bucket YOUR_S3_BUCKET

sam deploy \
    --template-file pkg/packaged-template.json \
    --stack-name sam-with-sfn \
    --capabilities CAPABILITY_IAM
```