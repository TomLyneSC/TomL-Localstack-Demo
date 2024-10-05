package localstack

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// localstack/AWS region
const (
	Region = "eu-west-2"
)

// getURL gets the required localstack url
func getURL() string {
	return "http://localhost:4566"
}

// CreateSession creates an AWS session based on Localstack config
func CreateSession() (sess *session.Session) {
	cfg := aws.NewConfig().
		WithRegion(Region).
		WithEndpoint(getURL()).
		WithCredentials((credentials.NewStaticCredentials("test", "test", "")))
	sess = session.Must(session.NewSession(cfg))
	return
}
