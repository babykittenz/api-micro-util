package dynamodb

import (
	toolkit "github.com/babykittenz/api-micro-util"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTraineeDDBRepository(t *testing.T) {
	// set up test env
	toolkit.SetupTest(t)
	// make me a fake trainee based off models.trainee
	fakeTrainee := models.Trainee{
		ID:                    "1",
		Name:                  "test",
		Email:                 "mike@q4impact.com",
		FirstName:             "mike",
		LastName:              "test",
		Phone:                 "1234567890",
		Company:               "q4impact",
		CompanyName:           "Q4 Impact",
		VisitorType:           "guest",
		MSHA:                  "",
		TruckNumber:           "",
		PreferredLanguage:     "en",
		LastTraining:          "12314325345",
		LastTrainingVideo:     "12314325345",
		LastTrainingAgreement: "12314325345",
		CompanyID:             "1231",
		LocationID:            "123143",
		RegionID:              "12314",
		CheckedIn:             false,
	}

	// Get the mock client
	client := toolkit.SafeGetDynamoDBClient().(*toolkit.MockDynamoDBClient)

	// Initialize the test repository
	repo := NewTestTraineeDDBRepository(client)

	err := repo.Save(&fakeTrainee)
	if err != nil {
		t.Error(err)
	}
	var trainee *models.Trainee
	trainee, err = repo.FindByID("2")

	if err != nil {
		t.Error(err)
	}
	log.Println(trainee)
	assert.Equal(t, "2", trainee.ID)

	err = repo.Delete("2")
	if err != nil {
		t.Error(err)
	}

	trainee, err = repo.FindByID("2")
	log.Println(trainee)
	assert.Nil(t, trainee)

}
