package localstack

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// DynamoClient implementation of DynamoDB service that allows for interactions with DynamoDB
type DynamoClient struct {
	dynamo dynamodbiface.DynamoDBAPI
}

// NewDynamoClient builds an instance of a dynamo service
func NewDynamoClient() (*DynamoClient, error) {
	return &DynamoClient{
		dynamo: dynamodb.New(CreateSession()),
	}, nil
}

// PutItem saves record to DynamoDB
func (d *DynamoClient) PutItem(input dynamodb.PutItemInput) error {
	_, err := d.dynamo.PutItem(&input)

	if err != nil {
		return err
	}
	return nil
}

// GetItem Queries DynamoDB
func (d *DynamoClient) GetItem(getItemInput dynamodb.GetItemInput) (map[string]*dynamodb.AttributeValue, error) {
	resp, err := d.dynamo.GetItem(&getItemInput)
	if err != nil {
		return nil, err
	}

	return resp.Item, nil
}

// UpdateItem updates record within DynamoDB
func (d *DynamoClient) UpdateItem(updateItemInput dynamodb.UpdateItemInput) error {
	_, err := d.dynamo.UpdateItem(&updateItemInput)

	if err != nil {
		return err
	}
	return nil
}

// DeleteItem deletes record from DynamoDB
func (d *DynamoClient) DeleteItem(deleteItemInput dynamodb.DeleteItemInput) error {
	_, err := d.dynamo.DeleteItem(&deleteItemInput)

	if err != nil {
		return err
	}
	return nil
}

// Query Queries DynamoDB
func (d *DynamoClient) Query(queryInput dynamodb.QueryInput) ([]map[string]*dynamodb.AttributeValue, error) {
	resp, err := d.dynamo.Query(&queryInput)
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}
