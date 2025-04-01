package toolkit

import (
	"context"
	"testing"

	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

// Test the mock implementation of GetItem
func TestMockGetItem(t *testing.T) {
	SetupTest(t)

	// Create a context for the request
	ctx := context.Background()

	// Get the mock client
	client := SafeGetDynamoDBClient().(*MockDynamoDBClient)

	// Create params for an existing trainee
	params := &ddb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "1"},
		},
	}

	// Call GetItem
	result, err := client.GetItem(ctx, params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Item)

	// Check for a non-existent item
	params = &ddb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "999"},
		},
	}

	result, err = client.GetItem(ctx, params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result.Item)
}
