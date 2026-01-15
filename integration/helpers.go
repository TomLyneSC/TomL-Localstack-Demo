package helpers

import (
	"bytes"
	"localstack-demo/src/localstack"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func AddBannedPlayer(name string, db *localstack.DynamoClient) {
	attribute := types.AttributeValueMemberS{
		Value: name,
	}

	item := dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"Name": &attribute,
		},
		TableName: aws.String("BannedPlayers"),
	}
	db.PutItem(item)
}

func RemoveBannedPlayer(name string, db *localstack.DynamoClient) {
	attribute := types.AttributeValueMemberS{
		Value: name,
	}

	item := dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
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
