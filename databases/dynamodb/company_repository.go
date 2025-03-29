package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// CompanyDDBRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on company records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type CompanyDDBRepository struct {
	client    *dynamodb.Client
	tableName string
}

// NewCompanyDDBRepository creates a new instance of a CompanyRepository using a DynamoDB client and a predefined table name.
func NewCompanyDDBRepository(client *dynamodb.Client) repository.CompanyRepository {
	return &CompanyDDBRepository{
		client:    client,
		tableName: "companies",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier and returns the result.
func (r *CompanyDDBRepository) FindByID(id string) (*models.Company, error) {
	return nil, nil
}

// FindAll retrieves all company records from the DynamoDB table and returns them as a slice of Company pointers.
func (r *CompanyDDBRepository) FindAll() ([]*models.Company, error) {
	return nil, nil
}

// Save stores or inserts the given company record into the DynamoDB table. Returns an error if the operation fails.
func (r *CompanyDDBRepository) Save(company *models.Company) error {
	return nil
}

// Update modifies an existing company record in the DynamoDB table and returns an error if the operation fails.
func (r *CompanyDDBRepository) Update(company *models.Company) error {
	return nil
}

// Delete removes a company record from the DynamoDB table based on the provided unique identifier and returns an error if it fails.
func (r *CompanyDDBRepository) Delete(id string) error {
	return nil
}
