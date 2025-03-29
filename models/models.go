package models

import "time"

// User represents a system user with authentication and access permissions.
type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	CompanyID  string `json:"company_id"`
	RegionID   string `json:"region_id"`
	LocationID string `json:"location_id"`
}

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

// Training represents a record of a training session completed by a trainee.
type Training struct {
	ID             string    `json:"id"`
	PDF            string    `json:"pdf,omitempty"`
	DateCompleted  time.Time `json:"date_completed"`
	ManualAddition bool      `json:"manual_addition,omitempty"`
	TraineeID      string    `json:"trainee_id,omitempty"`
	LocationID     string    `json:"location_id,omitempty"`
	RegionID       string    `json:"region_id,omitempty"`
	CompanyID      string    `json:"company_id,omitempty"`
}

// Checkin represents a record of when a trainee checks in to a location.
type Checkin struct {
	ID         string    `json:"id"`
	TraineeID  string    `json:"trainee_id,omitempty"`
	LocationID string    `json:"location_id,omitempty"`
	RegionID   string    `json:"region_id,omitempty"`
	CompanyID  string    `json:"company_id,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	Type       string    `json:"type,omitempty"`
}

// Company represents the details of a company within the system.
type Company struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Address           string `json:"address"`
	City              string `json:"city"`
	State             string `json:"state"`
	Zip               string `json:"zip"`
	StorageName       string `json:"storage_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Logo              string `json:"logo"`
	BlankPDF          string `json:"blank_pdf"`
	Map               string `json:"map"`
	Description       string `json:"description"`
	PreferredLanguage string `json:"preferred_language"`
	HasAgreement      bool   `json:"has_agreement"`
	Agreement         string `json:"agreement,omitempty"`
	HasVideo          bool   `json:"has_video"`
	Video             string `json:"video,omitempty"`
}

// Region represents a geographic or operational region associated with a company.
type Region struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	Phone             string   `json:"phone"`
	CompanyID         string   `json:"company_id"`
	StorageName       string   `json:"storage_name"`
	Address           string   `json:"address"`
	City              string   `json:"city"`
	State             string   `json:"state"`
	Zip               string   `json:"zip"`
	Map               string   `json:"map"`
	Logo              string   `json:"logo"`
	BlankPDF          string   `json:"blank_pdf"`
	Description       string   `json:"description,omitempty"`
	PreferredLanguage string   `json:"preferred_lang"`
	HasAgreement      bool     `json:"has_agreement"`
	Agreement         string   `json:"agreement,omitempty"`
	HasVideo          bool     `json:"has_video"`
	Video             string   `json:"video,omitempty"`
	Languages         []string `json:"languages"`
}

// Location represents a specific site or facility within a region.
type Location struct {
	ID                             string   `json:"id"`
	CompanyID                      string   `json:"company_id"`
	RegionID                       string   `json:"region_id"`
	Name                           string   `json:"name"`
	StorageName                    string   `json:"storage_name"`
	Phone                          string   `json:"phone"`
	Email                          string   `json:"email"`
	Address                        string   `json:"address"`
	State                          string   `json:"state"`
	City                           string   `json:"city"`
	Zip                            string   `json:"zip"`
	CheckinTextMessages            bool     `json:"checkin_text_messages"`
	TextNotificationsNumber        string   `json:"text_notifications_number,omitempty"`
	Map                            string   `json:"map"`
	Logo                           string   `json:"logo"`
	BlankPDF                       string   `json:"blank_pdf"`
	Description                    string   `json:"description,omitempty"`
	PreferredLanguage              string   `json:"preferred_lang"`
	HasAgreement                   bool     `json:"has_agreement"`
	Agreement                      string   `json:"agreement,omitempty"`
	ExpirationTime                 string   `json:"expiration_time"`
	HasVideo                       bool     `json:"has_video"`
	Video                          string   `json:"video,omitempty"`
	Languages                      []string `json:"languages"`
	MshaID                         string   `json:"msha_id,omitempty"`
	DateOfTrainingPlaceholder      string   `json:"date_of_training_placeholder"`
	TraineeNamePlaceholder         string   `json:"trainee_name_placeholder"`
	PhonePlaceholder               string   `json:"phone_placeholder"`
	TrainingLocationPlaceholder    string   `json:"training_location_placeholder"`
	EmailPlaceholder               string   `json:"email_placeholder"`
	MshaNumberPlaceholder          string   `json:"msha_number_placeholder"`
	ScalehouseAttendantPlaceholder string   `json:"scalehouse_attendant_placeholder"`
	TruckNumberPlaceholder         string   `json:"truck_number_placeholder"`
	TrainingPerformedPlaceholder   string   `json:"training_performed_placeholder"`
	CompanyPlaceholder             string   `json:"company_placeholder"`
}

// Language represents all the text strings used for UI localization and customization.
type Language struct {
	WelcomeText                   string `json:"welcome_text"`
	WelcomeText2                  string `json:"welcome_text_2"`
	BottomText                    string `json:"bottom_text"`
	StartTrainingButtonText       string `json:"start_training_button_text"`
	WhosTrainingWelcomeText       string `json:"whos_training_welcome_text"`
	WhosTrainingWelcomeSubtext    string `json:"whos_training_welcome_subtext"`
	FormFirstName                 string `json:"form_first_name"`
	FormFirstNamePlaceholder      string `json:"form_first_name_placeholder"`
	FormLastName                  string `json:"form_last_name"`
	FormLastNamePlaceholder       string `json:"form_last_name_placeholder"`
	FormEmail                     string `json:"form_email"`
	FormEmailPlaceholder          string `json:"form_email_placeholder"`
	FormCompany                   string `json:"form_company"`
	FormCompanyPlaceholder        string `json:"form_company_placeholder"`
	FormPhone                     string `json:"form_phone"`
	FormPhonePlaceholder          string `json:"form_phone_placeholder"`
	FormVisitorType               string `json:"form_visitor_type"`
	FormVisitorTypePleaseSelect   string `json:"form_visitor_type_please_select"`
	FormVisitorTypeContractor     string `json:"form_visitor_type_contractor"`
	FormVisitorTypeVendor         string `json:"form_visitor_type_vendor"`
	FormVisitorTypeGuest          string `json:"form_visitor_type_guest"`
	FormMshaNumber                string `json:"form_msha_number"`
	FormMshaNumberPlaceholder     string `json:"form_msha_number_placeholder"`
	FormTruckNumber               string `json:"form_truck_number"`
	FormTruckNumberPlaceholder    string `json:"form_truck_number_placeholder"`
	FormPreferredLang             string `json:"form_preferred_lang"`
	FormPreferredLangPleaseSelect string `json:"form_preferred_lang_please_select"`
	FormPreferredLangEn           string `json:"form_preferred_lang_en"`
	FormPreferredLangEs           string `json:"form_preferred_lang_es"`
	AddTraineeButton              string `json:"add_trainee_button"`
	GoBackButton                  string `json:"go_back_button"`
	ContinueButton                string `json:"continue_button"`
	FormNewHereText               string `json:"form_new_here_text"`
	FormNewHereSubtext            string `json:"form_new_here_subtext"`
	FormWelcomeBackText           string `json:"form_welcome_back_text"`
	FormWelcomeBackSubtext        string `json:"form_welcome_back_subtext"`
	FormTrainingExpiredText       string `json:"form_training_expired_text"`
	FormTrainingWillExpireText    string `json:"form_training_will_expire_text"`
	IsThisYouText                 string `json:"is_this_you_text"`
	IsThisYouSubtext              string `json:"is_this_you_subtext"`
	NotMeButton                   string `json:"not_me_button"`
	ThisIsMeButton                string `json:"this_is_me_button"`
	AnyOfTheseYouText             string `json:"any_of_these_you_text"`
	AnyOfTheseYouSubtext          string `json:"any_of_these_you_subtext"`
	NoneAreMeButton               string `json:"none_are_me_button"`
	RestartVideo                  string `json:"restart_video"`
	StartVideo                    string `json:"start_video"`
	TraineeListText               string `json:"trainee_list_text"`
	TraineeListSubtext            string `json:"trainee_list_subtext"`
	StartTrainingButton           string `json:"start_training_button"`
	AgreementCheckboxText         string `json:"agreement_checkbox_text"`
	SignatureSubtext              string `json:"signature_subtext"`
	ClearButton                   string `json:"clear_button"`
	YourDoneText                  string `json:"your_done_text"`
	YourDoneSubtext               string `json:"your_done_subtext"`
	YourDoneSubtextCountdown      string `json:"your_done_subtext_countdown"` // there is a typo in the original file
}

// TextMessage represents a text message template that can be sent to recipients.
type TextMessage struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Text       string    `json:"text"`
	Type       string    `json:"type"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	LocationID string    `json:"location_id"`
}

// AutomaticTextMessage represents a scheduled text message to be sent to recipients.
type AutomaticTextMessage struct {
	ID                 string   `json:"id"`
	RuleName           string   `json:"ruleName"`
	ScheduleExpression string   `json:"scheduleExpression"`
	Message            string   `json:"message"`
	Recipients         []string `json:"recipients"`
	Title              string   `json:"title"`
	Active             bool     `json:"active"`
	LocationID         string   `json:"location_id"`
	DayOfWeek          string   `json:"dayOfWeek"`
	DayOfMonth         string   `json:"dayOfMonth"`
	Frequency          string   `json:"frequency"`
	TimeToSend         string   `json:"timeToSend"`
	RecipientType      string   `json:"recipientType"`
	MessageID          string   `json:"message_id"`
}
