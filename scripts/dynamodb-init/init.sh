#!/bin/bash
set -e

# wait for dynamodb to start
echo "Waiting for DynamoDB to start..."
sleep 5

export AWS_ACCESS_KEY_ID=dummy
export AWS_SECRET_ACCESS_KEY=dummy
export AWS_DEFAULT_REGION=ap-northeast-1

echo "Creating DynamoDB tables..."
aws dynamodb create-table \
  --endpoint-url http://dynamodb:8000 \
  --cli-input-json file:///init/video-metadata.json

echo "DynamoDB initialization completed"
