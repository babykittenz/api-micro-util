package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoTextMessageRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on TextMessage records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoTextMessageRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewTextMessageRepository(client *dynamodb.Client) repository.TextMessageRepository {
	return &DynamoTextMessageRepository{
		client:    client,
		tableName: "text_messages",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoTextMessageRepository) FindByID(id string) (*models.TextMessage, error) {
	return nil, nil
}

func (r *DynamoTextMessageRepository) FindAll() ([]*models.TextMessage, error) {
	return nil, nil
}

func (r *DynamoTextMessageRepository) FindAllByLocationID(id string) ([]*models.TextMessage, error) {
	return nil, nil
}

func (r *DynamoTextMessageRepository) Save(textMessage *models.TextMessage) error {
	return nil
}

func (r *DynamoTextMessageRepository) Update(textMessage *models.TextMessage) error {
	return nil
}

func (r *DynamoTextMessageRepository) Delete(id string) error {
	return nil
}
