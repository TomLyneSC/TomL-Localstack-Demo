#!/usr/bin/env bash
set -o pipefail -ex

# start localstack container
echo "Starting localstack"
docker-compose -f docker-compose.yaml up -d --build localstack

echo "Wait"
sleep 10

# make aws request point to localstack
echo "Setup localstack aws resources"
ENDPOINT_URL="http://localhost:4566"

# create dynamo table locally
echo “Adding Mock Dynamo Table” 
aws --endpoint-url $ENDPOINT_URL --region=eu-west-2 dynamodb create-table --table-name BannedPlayers --attribute-definitions AttributeName=Name,AttributeType=S --key-schema AttributeName=Name,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

# build service
CGO_ENABLED=0 go build -o demo ./src/main.go