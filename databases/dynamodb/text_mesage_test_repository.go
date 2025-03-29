package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

type testTextMessageDDBRepository struct {
	client *dynamodb.Client
}

func NewTestTextMessageDDBRepository(client *dynamodb.Client) repository.TextMessageRepository {
	return &testTextMessageDDBRepository{
		client: nil,
	}
}

// FindByID retrieves a TextMessage record from the DynamoDB table using the provided unique identifier.
func (r *testTextMessageDDBRepository) FindByID(id string) (*models.TextMessage, error) {
	return nil, nil
}

// FindAll retrieves all TextMessage records from the DynamoDB table and returns them as a slice.
func (r *testTextMessageDDBRepository) FindAll() ([]*models.TextMessage, error) {
	return nil, nil
}

// FindAllByLocationID retrieves all TextMessage records associated with the specified LocationID from the DynamoDB table.
func (r *testTextMessageDDBRepository) FindAllByLocationID(id string) ([]*models.TextMessage, error) {
	return nil, nil
}

// Save stores the given TextMessage in the DynamoDB table and returns an error if the operation fails.
func (r *testTextMessageDDBRepository) Save(textMessage *models.TextMessage) error {
	return nil
}

// Update modifies an existing TextMessage record in the DynamoDB table and returns an error if the operation fails.
func (r *testTextMessageDDBRepository) Update(textMessage *models.TextMessage) error {
	return nil
}

// Delete removes a TextMessage record from the DynamoDB table by its unique identifier and returns an error if it fails.
func (r *testTextMessageDDBRepository) Delete(id string) error {
	return nil
}
