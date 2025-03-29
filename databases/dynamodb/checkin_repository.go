package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoCheckinRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on Checkin records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoCheckinRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewCheckinRepository(client *dynamodb.Client) repository.CheckinRepository {
	return &DynamoCheckinRepository{
		client:    client,
		tableName: "checkins",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoCheckinRepository) FindByID(id string) (*models.Checkin, error) {
	return nil, nil
}

func (r *DynamoCheckinRepository) FindAll() ([]*models.Checkin, error) {
	return nil, nil
}

func (r *DynamoCheckinRepository) FindAllByCompanyID(id string) ([]*models.Checkin, error) {
	return nil, nil
}

func (r *DynamoCheckinRepository) FindAllByRegionID(id string) ([]*models.Checkin, error) {
	return nil, nil
}

func (r *DynamoCheckinRepository) FindAllByLocationID(id string) ([]*models.Checkin, error) {
	return nil, nil
}

func (r *DynamoCheckinRepository) FindAllByTraineeID(id string) ([]*models.Checkin, error) {
	return nil, nil
}

func (r *DynamoCheckinRepository) Save(checkin *models.Checkin) error {
	return nil
}

func (r *DynamoCheckinRepository) Update(checkin *models.Checkin) error {
	return nil
}

func (r *DynamoCheckinRepository) Delete(id string) error {
	return nil
}
