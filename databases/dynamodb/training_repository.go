package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

// DynamoTrainingRepository is a repository implementation for managing company data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on Training records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type DynamoTrainingRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewTrainingRepository(client *dynamodb.Client) repository.TrainingRepository {
	return &DynamoTrainingRepository{
		client:    client,
		tableName: "trainings",
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier.
func (r *DynamoTrainingRepository) FindByID(id string) (*models.Training, error) {
	return nil, nil
}

func (r *DynamoTrainingRepository) FindAll() ([]*models.Training, error) {
	return nil, nil
}

func (r *DynamoTrainingRepository) FindAllByCompanyID(id string) ([]*models.Training, error) {
	return nil, nil
}

func (r *DynamoTrainingRepository) FindAllByRegionID(id string) ([]*models.Training, error) {
	return nil, nil
}

func (r *DynamoTrainingRepository) FindAllByLocationID(id string) ([]*models.Training, error) {
	return nil, nil
}

func (r *DynamoTrainingRepository) FindAllByTraineeID(id string) ([]*models.Training, error) {
	return nil, nil
}

func (r *DynamoTrainingRepository) Save(training *models.Training) error {
	return nil
}

func (r *DynamoTrainingRepository) Update(training *models.Training) error {
	return nil
}

func (r *DynamoTrainingRepository) Delete(id string) error {
	return nil
}
