ls .
aws s3 sync s3://$ENV_BUCKET .
cat ./serverless.env.yml
serverless deploy --stage=dev