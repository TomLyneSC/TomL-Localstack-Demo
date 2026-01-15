package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// BuildAWSConfig takes the default config and enriches it with region and endpoint information
func BuildAWSConfig(region string, endpoint string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return aws.Config{}, err
	}

	if endpoint != "" {
		cfg.BaseEndpoint = aws.String(endpoint)
		cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           endpoint,
					SigningRegion: region,
				}, nil
			})
	}

	return cfg, nil
}

func main() {
	// These two lines would never be past of a go program normally and would be passed in via config options
	// But for this demo, we are going to hardcode them here
	// Don't comment out these lines or else it will try to talk to the real AWS
	region := "eu-west-2"
	endpoint := "http://localhost:4566"

	// This is the REAL start of the application

	// Setup AWS connection - first, let's grab the config stored in the secret credentials file
	cfg, err := BuildAWSConfig(region, endpoint)
	if err != nil {
		return
	}

	// Using the config/credentials, let's setup the service that'll let us talk to dynamoDB
	dynamoService := dynamodb.NewFromConfig(cfg)

	// Read command line args
	targetName := flag.CommandLine.String("targetName", "", "Test")
	flag.Parse()
	if targetName == nil || len(*targetName) == 0 {
		return
	}

	// Scan DynamoDB table for list of banned players
	in := dynamodb.ScanInput{
		TableName: aws.String("BannedPlayers"),
	}
	out, _ := dynamoService.Scan(context.TODO(), &in)
	items := out.Items

	// Loop through output to find the banned player
	foundTarget := false
	for _, item := range items {
		if nameAttr, ok := item["Name"]; ok && nameAttr != nil {
			if sAttr, ok := nameAttr.(*types.AttributeValueMemberS); ok && sAttr.Value == *targetName {
				foundTarget = true
				break
			}
		}
	}

	// Return if we found the banned player
	fmt.Print(foundTarget)
}
