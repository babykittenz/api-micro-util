package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoRegionRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on Region records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoRegionRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewRegionRepository(client *dynamodb.Client) repository.RegionRepository {
	return &DynamoRegionRepository{
		client:    client,
		tableName: "regions",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoRegionRepository) FindByID(id string) (*models.Region, error) {
	return nil, nil
}

func (r *DynamoRegionRepository) FindAll() ([]*models.Region, error) {
	return nil, nil
}

func (r *DynamoRegionRepository) FindAllByCompanyID(id string) ([]*models.Region, error) {
	return nil, nil
}

func (r *DynamoRegionRepository) Save(region *models.Region) error {
	return nil
}

func (r *DynamoRegionRepository) Update(region *models.Region) error {
	return nil
}

func (r *DynamoRegionRepository) Delete(id string) error {
	return nil
}
