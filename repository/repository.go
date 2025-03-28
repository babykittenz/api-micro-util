package repository

import "github.com/babykittenz/api-micro-util/models"

type TraineeRepository interface {
	FindByID(id string) (*models.Trainee, error)
	FindByEmail(email string) (*models.Trainee, error)
	Save(trainee *models.Trainee) error
	Update(trainee *models.Trainee) error
	Delete(id string) error
}

type CompanyRepository interface {
	FindByID(id string) (*models.Company, error)
	FindAll() ([]*models.Company, error)
	Save(company *models.Company) error
}
