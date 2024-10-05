package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewConfig proxy underlying aws package
var NewConfig = aws.NewConfig

// BuildAWSConfig generate a standard AWS Config, passing in endpoint and region for CI Integration tests
func BuildAWSConfig(region string, endpoint string) aws.Config {
	config := aws.NewConfig().WithRegion(region)
	if endpoint != "" {
		config.WithEndpoint(endpoint)
	}

	return *config
}

// CreateSession creates a new AWS session.
func CreateSession(region string) (sess *session.Session, err error) {
	// Create AWS Config.
	if len(region) == 0 {
		return nil, fmt.Errorf("no valid AWS_REGION found")
	}
	cfg := NewConfig().WithRegion(region)

	// Create AWS Session.
	sess, err = session.NewSession(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS session")
	}

	return
}

func main() {
	// These two lines would never be past of a go program normally and would be passed in via config options
	// But for this demo, we are going to hardcode them here
	// Don't comment out these lines or else it will try to talk to the real AWS
	region := "eu-west-2"
	endpoint := "http://localhost:4566"

	// This is the start of the application
	awsConfig := BuildAWSConfig(region, endpoint)

	// Setup AWS connection
	sess, err := CreateSession("eu-west-2")
	if err != nil {
		return
	}
	dynamoService := dynamodb.New(sess, &awsConfig)

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
	out, _ := dynamoService.Scan(&in)
	items := out.Items

	// Loop through list to find banned player
	foundTarget := false
	for _, item := range items {
		if *item["Name"].S == *targetName {
			foundTarget = true
			break
		}
	}

	// Return if we found the banned player
	fmt.Print(foundTarget)
}
