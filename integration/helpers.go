package helpers

import (
	"bytes"
	"localstack-demo/src/localstack"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func AddBannedPlayer(name string, db *localstack.DynamoClient) {
	attribute := dynamodb.AttributeValue{
		S: aws.String(name),
	}

	item := dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Name": &attribute,
		},
		TableName: aws.String("BannedPlayers"),
	}
	db.PutItem(item)
}

func RemoveBannedPlayer(name string, db *localstack.DynamoClient) {
	attribute := dynamodb.AttributeValue{
		S: aws.String(name),
	}

	item := dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Name": &attribute,
		},
		TableName: aws.String("BannedPlayers"),
	}
	db.DeleteItem(item)
}

func RunBannedCheck(targetName string) string {
	cmd := exec.Command("../../demo", "-targetName", targetName)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	cmd.Run()
	return outb.String()
}
