#!/bin/bash
set -e

# wait for s3 to start
echo "Waiting for S3 to start..."
sleep 5

echo "Creating S3 bucket..."
aws s3api create-bucket --bucket local-video-bucket --endpoint-url http://minio:9000
