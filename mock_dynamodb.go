package toolkit

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Sample trainees data for tests
var testTrainees = []map[string]interface{}{
	{
		"id":       "1",
		"name":     "John Doe",
		"email":    "john@example.com",
		"category": "developer",
	},
	{
		"id":       "2",
		"name":     "Jane Smith",
		"email":    "jane@example.com",
		"category": "designer",
	},
}

// MockDynamoDBClient implements all required DynamoDB methods for testing
type MockDynamoDBClient struct {
	t *testing.T // for test assertions if needed
}

// GetItem implementation for mock
func (m *MockDynamoDBClient) GetItem(ctx context.Context, params *ddb.GetItemInput, optFns ...func(*ddb.Options)) (*ddb.GetItemOutput, error) {
	id := params.Key["id"].(*types.AttributeValueMemberS).Value

	// Look for matching trainee
	for _, trainee := range testTrainees {
		if trainee["id"] == id {
			item, err := attributevalue.MarshalMap(trainee)
			if err != nil {
				return nil, err
			}
			return &ddb.GetItemOutput{Item: item}, nil
		}
	}

	// Return empty if not found
	return &ddb.GetItemOutput{}, nil
}

// Scan implementation for mock
func (m *MockDynamoDBClient) Scan(ctx context.Context, params *ddb.ScanInput, optFns ...func(*ddb.Options)) (*ddb.ScanOutput, error) {
	items := make([]map[string]types.AttributeValue, 0, len(testTrainees))
	for _, trainee := range testTrainees {
		item, err := attributevalue.MarshalMap(trainee)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &ddb.ScanOutput{
		Items: items,
		Count: int32(len(items)),
	}, nil
}

// PutItem implementation for mock
func (m *MockDynamoDBClient) PutItem(ctx context.Context, params *ddb.PutItemInput, optFns ...func(*ddb.Options)) (*ddb.PutItemOutput, error) {
	// Just return success
	return &ddb.PutItemOutput{}, nil
}

// DeleteItem implementation for mock
func (m *MockDynamoDBClient) DeleteItem(ctx context.Context, params *ddb.DeleteItemInput, optFns ...func(*ddb.Options)) (*ddb.DeleteItemOutput, error) {
	// Just return success
	return &ddb.DeleteItemOutput{}, nil
}

// UpdateItem implementation for mock
func (m *MockDynamoDBClient) UpdateItem(ctx context.Context, params *ddb.UpdateItemInput, optFns ...func(*ddb.Options)) (*ddb.UpdateItemOutput, error) {
	// Just return success
	return &ddb.UpdateItemOutput{}, nil
}

// Query implementation for mock
func (m *MockDynamoDBClient) Query(ctx context.Context, params *ddb.QueryInput, optFns ...func(*ddb.Options)) (*ddb.QueryOutput, error) {
	// Simple implementation that returns the test trainees that match the category filter
	category := params.ExpressionAttributeValues[":category"].(*types.AttributeValueMemberS).Value
	items := make([]map[string]types.AttributeValue, 0)

	for _, trainee := range testTrainees {
		if traineeCategory, ok := trainee["category"].(string); ok && traineeCategory == category {
			item, err := attributevalue.MarshalMap(trainee)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
	}

	return &ddb.QueryOutput{
		Items: items,
		Count: int32(len(items)),
	}, nil
}
