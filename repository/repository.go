package repository

import "github.com/babykittenz/api-micro-util/models"

// TraineeRepository provides methods to perform CRUD operations and specific queries on Trainee data.
type TraineeRepository interface {
	FindByID(id string) (*models.Trainee, error)
	FindByEmail(email string) (*models.Trainee, error)
	FindByPhone(phone string) (*models.Trainee, error)
	FindByPhoneAndLocation(phone string, location string) (*models.Trainee, error)
	FindByEmailAndLocation(email string, location string) (*models.Trainee, error)
	FindByNames(firstName string, lastName string) (*models.Trainee, error)
	FindByNamesAndLocation(firstName string, lastName string, location string) (*models.Trainee, error)
	FindAll() ([]*models.Trainee, error)
	FindAllByCompanyID(id string) ([]*models.Trainee, error)
	FindAllByRegionID(id string) ([]*models.Trainee, error)
	FindAllByLocationID(id string) ([]*models.Trainee, error)
	Save(trainee *models.Trainee) error
	Update(trainee *models.Trainee) error
	Delete(id string) error
	DeleteByEmail(email string) error
	CompleteTraining(id string) error
	Checkin(id string) error
	Checkout(id string) error
}

// TrainingRepository defines the interface for interacting with Training data storage and management operations.
type TrainingRepository interface {
	FindByID(id string) (*models.Training, error)
	FindAll() ([]*models.Training, error)
	FindAllByCompanyID(id string) ([]*models.Training, error)
	FindAllByRegionID(id string) ([]*models.Training, error)
	FindAllByLocationID(id string) ([]*models.Training, error)
	FindAllByTraineeID(id string) ([]*models.Training, error)
	Save(training *models.Training) error
	Update(training *models.Training) error
	Delete(id string) error
}

// CheckinRepository defines a contract for managing check-in records in a storage system.
// It provides methods for retrieving, creating, updating, and deleting check-in data.
// The repository operates on models.Checkin entities for CRUD operations.
type CheckinRepository interface {
	FindByID(id string) (*models.Checkin, error)
	FindAll() ([]*models.Checkin, error)
	FindAllByCompanyID(id string) ([]*models.Checkin, error)
	FindAllByRegionID(id string) ([]*models.Checkin, error)
	FindAllByTraineeID(id string) ([]*models.Checkin, error)
	Save(checkin *models.Checkin) error
	Update(checkin *models.Checkin) error
	Delete(id string) error
}

// CompanyRepository defines the methods required for managing company data within a persistence layer.
type CompanyRepository interface {
	FindByID(id string) (*models.Company, error)
	FindAll() ([]*models.Company, error)
	Save(company *models.Company) error
	Update(company *models.Company) error
	Delete(id string) error
}

// RegionRepository defines the interface for managing operations related to regions in the repository.
// Provides methods to find, save, update, and delete regions.
type RegionRepository interface {
	FindByID(id string) (*models.Region, error)
	FindAll() ([]*models.Region, error)
	FindAllByCompanyID(id string) ([]*models.Region, error)
	Save(region *models.Region) error
	Update(region *models.Region) error
	Delete(id string) error
}

// LocationRepository defines the interface for managing and executing operations on Location entities.
// It provides methods to retrieve, save, update, and delete location data.
type LocationRepository interface {
	FindByID(id string) (*models.Location, error)
	FindAll() ([]*models.Location, error)
	FindAllByCompanyID(id string) ([]*models.Location, error)
	FindAllByRegionID(id string) ([]*models.Location, error)
	Save(location *models.Location) error
	Update(location *models.Location) error
	Delete(id string) error
}

// LanguageRepository defines methods to manage Language resources in storage.
type LanguageRepository interface {
	FindByID(id string) (*models.Language, error)
	FindAll() ([]*models.Language, error)
	Save(language *models.Language) error
	Update(language *models.Language) error
	Delete(id string) error
}

// TextMessageRepository defines methods to interact with and manage TextMessage data storage and retrieval.
type TextMessageRepository interface {
	FindByID(id string) (*models.TextMessage, error)
	FindAll() ([]*models.TextMessage, error)
	FindAllByLocationID(id string) ([]*models.TextMessage, error)
	Save(textMessage *models.TextMessage) error
	Update(textMessage *models.TextMessage) error
	Delete(id string) error
}

// AutomaticTextMessageRepository defines the interface for managing AutomaticTextMessage data storage and retrieval.
type AutomaticTextMessageRepository interface {
	FindByID(id string) (*models.AutomaticTextMessage, error)
	FindAll() ([]*models.AutomaticTextMessage, error)
	FindAllByLocationID(id string) ([]*models.AutomaticTextMessage, error)
	Save(automaticTextMessage *models.AutomaticTextMessage) error
	Update(automaticTextMessage *models.AutomaticTextMessage) error
	Delete(id string) error
}
