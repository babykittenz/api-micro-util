package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoAutomaticTextMessageRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on AutomaticTextMessage records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoAutomaticTextMessageRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewAutomaticTextMessageRepository(client *dynamodb.Client) repository.AutomaticTextMessageRepository {
	return &DynamoAutomaticTextMessageRepository{
		client:    client,
		tableName: "automatic_text_messages",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoAutomaticTextMessageRepository) FindByID(id string) (*models.AutomaticTextMessage, error) {
	return nil, nil
}

func (r *DynamoAutomaticTextMessageRepository) FindAll() ([]*models.AutomaticTextMessage, error) {
	return nil, nil
}

func (r *DynamoAutomaticTextMessageRepository) FindAllByLocationID(id string) ([]*models.AutomaticTextMessage, error) {
	return nil, nil
}

func (r *DynamoAutomaticTextMessageRepository) Save(automaticTextMessage *models.AutomaticTextMessage) error {
	return nil
}

func (r *DynamoAutomaticTextMessageRepository) Update(automaticTextMessage *models.AutomaticTextMessage) error {
	return nil
}

func (r *DynamoAutomaticTextMessageRepository) Delete(id string) error {
	return nil
}
