aws s3 cp s3://$ENV_BUCKET . --recursive
serverless deploy --stage=prod