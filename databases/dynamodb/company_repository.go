package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoCompanyRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on company records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoCompanyRepository struct {
	client    *dynamodb.Client
	tableName string
}

// NewCompanyRepository creates a new instance of a CompanyRepository using a DynamoDB client and a predefined table name.
func NewCompanyRepository(client *dynamodb.Client) repository.CompanyRepository {
	return &DynamoCompanyRepository{
		client:    client,
		tableName: "companies",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoCompanyRepository) FindByID(id string) (*models.Company, error) {
	return nil, nil
}

// FindAll retrieves all company records from the DynamoDB table.
func (r *DynamoCompanyRepository) FindAll() ([]*models.Company, error) {
	return nil, nil
}

// Save stores the given company instance into the DynamoDB table or updates it if it already exists.
func (r *DynamoCompanyRepository) Save(company *models.Company) error {
	return nil
}
