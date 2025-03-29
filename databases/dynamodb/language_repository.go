package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoLanguageRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on Language records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoLanguageRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewLanguageRepository(client *dynamodb.Client) repository.LanguageRepository {
	return &DynamoLanguageRepository{
		client:    client,
		tableName: "languages",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoLanguageRepository) FindByID(id string) (*models.Language, error) {
	return nil, nil
}

func (r *DynamoLanguageRepository) FindAll() ([]*models.Language, error) {
	return nil, nil
}

func (r *DynamoLanguageRepository) Save(language *models.Language) error {
	return nil
}

func (r *DynamoLanguageRepository) Update(language *models.Language) error {
	return nil
}

func (r *DynamoLanguageRepository) Delete(id string) error {
	return nil
}
