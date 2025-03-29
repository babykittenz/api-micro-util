package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoLocationRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on Location records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoLocationRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewLocationRepository(client *dynamodb.Client) repository.LocationRepository {
	return &DynamoLocationRepository{
		client:    client,
		tableName: "locations",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoLocationRepository) FindByID(id string) (*models.Location, error) {
	return nil, nil
}

func (r *DynamoLocationRepository) FindAll() ([]*models.Location, error) {
	return nil, nil
}

func (r *DynamoLocationRepository) FindAllByCompanyID(id string) ([]*models.Location, error) {
	return nil, nil
}

func (r *DynamoLocationRepository) FindAllByRegionID(id string) ([]*models.Location, error) {
	return nil, nil
}

func (r *DynamoLocationRepository) Save(location *models.Location) error {
	return nil
}

func (r *DynamoLocationRepository) Update(location *models.Location) error {
	return nil
}

func (r *DynamoLocationRepository) Delete(id string) error {
	return nil
}
