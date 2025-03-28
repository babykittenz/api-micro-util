package models

// Trainee represents the details of an individual undergoing training within the system.
type Trainee struct {
	ID                    string `json:"id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Name                  string `json:"display_name"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone"`
	Company               string `json:"company"`
	CompanyName           string `json:"display_company"`
	VisitorType           string `json:"visitor_type"`
	MSHA                  string `json:"msha_number"`
	TruckNumber           string `json:"truck_number"`
	PreferredLanguage     string `json:"preferred_language"`
	LastTraining          string `json:"last_training"`
	LastTrainingVideo     string `json:"last_training_video"`
	LastTrainingAgreement string `json:"last_training_agreement"`
	CompanyID             string `json:"company_id"`
	LocationID            string `json:"location_id"`
	RegionID              string `json:"region_id"`
	CheckedIn             bool   `json:"checked_in"`
	CheckinID             string `json:"checkin_id"`
}

type User struct{}

type Training struct{}

type Location struct{}

type Region struct{}

type Company struct{}

type Language struct{}

type TextMessage struct{}

type AutomaticTextMessage struct{}

type Form struct{}
