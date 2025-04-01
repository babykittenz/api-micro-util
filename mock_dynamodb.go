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
		"id":                      "2",
		"first_name":              "Robert",
		"last_name":               "Martinez",
		"display_name":            "Robert Martinez",
		"email":                   "robert.martinez@example.com",
		"phone":                   "555-456-7890",
		"company":                 "ABC Construction",
		"display_company":         "ABC Construction Inc.",
		"visitor_type":            "contractor",
		"msha_number":             "MSHA123456",
		"truck_number":            "T-789",
		"preferred_language":      "en",
		"last_training":           "2024-02-15",
		"last_training_video":     "2024-02-15",
		"last_training_agreement": "2024-02-15",
		"company_id":              "comp-001",
		"location_id":             "loc-001",
		"region_id":               "reg-001",
		"checked_in":              true,
	},
	{
		"id":                      "3",
		"first_name":              "Jennifer",
		"last_name":               "Garcia",
		"display_name":            "Jennifer Garcia",
		"email":                   "jennifer.garcia@example.com",
		"phone":                   "555-567-8901",
		"company":                 "XYZ Logistics",
		"display_company":         "XYZ Logistics LLC",
		"visitor_type":            "vendor",
		"msha_number":             "MSHA789012",
		"truck_number":            "T-456",
		"preferred_language":      "es",
		"last_training":           "2024-01-20",
		"last_training_video":     "2024-01-20",
		"last_training_agreement": "2024-01-20",
		"company_id":              "comp-002",
		"location_id":             "loc-002",
		"region_id":               "reg-002",
		"checked_in":              false,
	},
	{
		"id":                      "4",
		"first_name":              "Thomas",
		"last_name":               "Wilson",
		"display_name":            "Thomas Wilson",
		"email":                   "thomas.wilson@example.com",
		"phone":                   "555-678-9012",
		"company":                 "123 Excavation",
		"display_company":         "123 Excavation Services",
		"visitor_type":            "contractor",
		"msha_number":             "MSHA345678",
		"truck_number":            "T-123",
		"preferred_language":      "en",
		"last_training":           "2024-03-05",
		"last_training_video":     "2024-03-05",
		"last_training_agreement": "2024-03-05",
		"company_id":              "comp-003",
		"location_id":             "loc-003",
		"region_id":               "reg-003",
		"checked_in":              true,
	},
}

// MockDynamoDBClient implements all required DynamoDB methods for testing
type MockDynamoDBClient struct {
	t           *testing.T // for test assertions if needed
	testingData []any
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
