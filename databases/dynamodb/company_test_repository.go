package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

type testCompanyDDBRepository struct {
	client *dynamodb.Client
}

func NewTestCompanyDDBRepository(client *dynamodb.Client) repository.CompanyRepository {
	return &testCompanyDDBRepository{
		client: nil,
	}
}

// FindByID retrieves a company record from the DynamoDB table by its unique identifier and returns the result.
func (r *testCompanyDDBRepository) FindByID(id string) (*models.Company, error) {
	return nil, nil
}

// FindAll retrieves all company records from the DynamoDB table and returns them as a slice of Company pointers.
func (r *testCompanyDDBRepository) FindAll() ([]*models.Company, error) {
	return nil, nil
}

// Save stores or inserts the given company record into the DynamoDB table. Returns an error if the operation fails.
func (r *testCompanyDDBRepository) Save(company *models.Company) error {
	return nil
}

// Update modifies an existing company record in the DynamoDB table and returns an error if the operation fails.
func (r *testCompanyDDBRepository) Update(company *models.Company) error {
	return nil
}

// Delete removes a company record from the DynamoDB table based on the provided unique identifier and returns an error if it fails.
func (r *testCompanyDDBRepository) Delete(id string) error {
	return nil
}
