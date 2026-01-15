package localstack

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoClient implementation of DynamoDB service that allows for interactions with DynamoDB
type DynamoClient struct {
	dynamo *dynamodb.Client
}

// NewDynamoClient builds an instance of a dynamo service
func NewDynamoClient() (*DynamoClient, error) {
	cfg, err := CreateConfig()
	if err != nil {
		return nil, err
	}
	return &DynamoClient{
		dynamo: dynamodb.NewFromConfig(cfg),
	}, nil
}

// PutItem saves record to DynamoDB
func (d *DynamoClient) PutItem(input dynamodb.PutItemInput) error {
	_, err := d.dynamo.PutItem(context.TODO(), &input)
	if err != nil {
		return err
	}
	return nil
}

// GetItem Queries DynamoDB
func (d *DynamoClient) GetItem(getItemInput dynamodb.GetItemInput) (map[string]types.AttributeValue, error) {
	resp, err := d.dynamo.GetItem(context.TODO(), &getItemInput)
	if err != nil {
		return nil, err
	}

	return resp.Item, nil
}

// UpdateItem updates record within DynamoDB
func (d *DynamoClient) UpdateItem(updateItemInput dynamodb.UpdateItemInput) error {
	_, err := d.dynamo.UpdateItem(context.TODO(), &updateItemInput)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem deletes record from DynamoDB
func (d *DynamoClient) DeleteItem(deleteItemInput dynamodb.DeleteItemInput) error {
	_, err := d.dynamo.DeleteItem(context.TODO(), &deleteItemInput)
	if err != nil {
		return err
	}
	return nil
}

// Query Queries DynamoDB
func (d *DynamoClient) Query(queryInput dynamodb.QueryInput) ([]map[string]types.AttributeValue, error) {
	resp, err := d.dynamo.Query(context.TODO(), &queryInput)
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}
