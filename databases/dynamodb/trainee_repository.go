package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
	"time"
)

// TraineeDDBRepository is a repository implementation for managing trainee data in DynamoDB.
// It utilizes a DynamoDB client and a specific table to perform CRUD operations on trainee records.
// The struct includes client for interaction with DynamoDB and tableName for specifying the target table.
type TraineeDDBRepository struct {
	client    *dynamodb.Client
	tableName string
}

// NewTraineeDDBRepository creates a new instance of a TraineeRepository using a DynamoDB client and a predefined table name.
func NewTraineeDDBRepository(client *dynamodb.Client, tableName string) repository.TraineeRepository {
	return &TraineeDDBRepository{
		client:    client,
		tableName: tableName,
	}
}

// FindByID retrieves a trainee record from DynamoDB based on the provided ID and returns a Trainee object or an error.
func (r *TraineeDDBRepository) FindByID(id string) (*models.Trainee, error) {
	ctx := context.Background()

	// Find the trainee by id
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get trainee from DynamoDB: %w", err)
	}

	// If no item found
	if result.Item == nil {
		return nil, fmt.Errorf("trainee with id %s not found", id)
	}

	// Unmarshal the DynamoDB item into a Trainee struct
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Item, &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByEmail retrieves a trainee record from DynamoDB based on the provided email and returns a Trainee object or an error.
func (r *TraineeDDBRepository) FindByEmail(email string) (*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by email
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with email %s not found", email)
	}

	// Unmarshal the first item (should be only one if email is unique)
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByPhone retrieves a trainee record from DynamoDB based on the provided phone number and returns a Trainee object or an error.
func (r *TraineeDDBRepository) FindByPhone(phone string) (*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by phone
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("phone = :phone"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":phone": &types.AttributeValueMemberS{Value: phone},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with phone %s not found", phone)
	}

	// Unmarshal the first item (should be only one if phone is unique)
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByPhoneAndLocation retrieves a trainee record from DynamoDB based on the provided phone and location ID. Returns a Trainee object or error.
func (r *TraineeDDBRepository) FindByPhoneAndLocation(phone string, locationID string) (*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by phone and locationID
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("phone = :phone AND location_id = :locationID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":phone":      &types.AttributeValueMemberS{Value: phone},
			":locationID": &types.AttributeValueMemberS{Value: locationID},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with phone %s and location_id %s not found", phone, locationID)
	}

	// Unmarshal the first item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByEmailAndLocation retrieves a trainee record from DynamoDB based on the provided email and location ID.
// Returns a Trainee object or an error if no matching record is found.
func (r *TraineeDDBRepository) FindByEmailAndLocation(email string, locationID string) (*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by email and locationID
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("email = :email AND location_id = :locationID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email":      &types.AttributeValueMemberS{Value: email},
			":locationID": &types.AttributeValueMemberS{Value: locationID},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with email %s and location_id %s not found", email, locationID)
	}

	// Unmarshal the first item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByNames retrieves a trainee record from DynamoDB based on the provided first and last names. Returns a Trainee object or error.
func (r *TraineeDDBRepository) FindByNames(firstName string, lastName string) (*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by first_name and last_name
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("first_name = :firstName AND last_name = :lastName"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":firstName": &types.AttributeValueMemberS{Value: firstName},
			":lastName":  &types.AttributeValueMemberS{Value: lastName},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with name %s %s not found", firstName, lastName)
	}

	// Unmarshal the first item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByNamesAndLocation retrieves a trainee record from DynamoDB using first name, last name, and location ID. Returns a Trainee object or an error.
func (r *TraineeDDBRepository) FindByNamesAndLocation(firstName string, lastName string, locationID string) (*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by first_name, last_name, and location_id
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("first_name = :firstName AND last_name = :lastName AND location_id = :locationID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":firstName":  &types.AttributeValueMemberS{Value: firstName},
			":lastName":   &types.AttributeValueMemberS{Value: lastName},
			":locationID": &types.AttributeValueMemberS{Value: locationID},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with name %s %s and location_id %s not found", firstName, lastName, locationID)
	}

	// Unmarshal the first item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindAllByCompanyID retrieves all trainee records associated with the given company ID from DynamoDB. Returns a slice of Trainee objects or an error.
func (r *TraineeDDBRepository) FindAllByCompanyID(id string) ([]*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by company_id
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("company_id = :companyID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":companyID": &types.AttributeValueMemberS{Value: id},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return []*models.Trainee{}, nil
	}

	// Unmarshal the items into a slice of Trainee objects
	var trainees []*models.Trainee
	err = attributevalue.UnmarshalListOfMaps(result.Items, &trainees)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainees: %w", err)
	}

	return trainees, nil
}

// FindAllByRegionID retrieves all trainee records associated with the specified region ID from DynamoDB.
// Returns a slice of Trainee objects or an error if the operation fails.
func (r *TraineeDDBRepository) FindAllByRegionID(id string) ([]*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by region_id
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("region_id = :regionID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":regionID": &types.AttributeValueMemberS{Value: id},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return []*models.Trainee{}, nil
	}

	// Unmarshal the items into a slice of Trainee objects
	var trainees []*models.Trainee
	err = attributevalue.UnmarshalListOfMaps(result.Items, &trainees)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainees: %w", err)
	}

	return trainees, nil
}

// FindAllByLocationID retrieves all trainee records associated with the specified location ID from DynamoDB.
// Returns a slice of Trainee objects or an error if the operation fails.
func (r *TraineeDDBRepository) FindAllByLocationID(id string) ([]*models.Trainee, error) {
	ctx := context.Background()

	// Define the scan parameters to filter by location_id
	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("location_id = :locationID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":locationID": &types.AttributeValueMemberS{Value: id},
		},
	}

	// Execute the scan operation
	result, err := r.client.Scan(ctx, scanInput)
	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Check if any items were found
	if len(result.Items) == 0 {
		return []*models.Trainee{}, nil
	}

	// Unmarshal the items into a slice of Trainee objects
	var trainees []*models.Trainee
	err = attributevalue.UnmarshalListOfMaps(result.Items, &trainees)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainees: %w", err)
	}

	return trainees, nil
}

// Save persists a trainee record to the DynamoDB table. Returns an error if the operation fails.
func (r *TraineeDDBRepository) Save(trainee *models.Trainee) error {
	ctx := context.Background()

	// Marshal the trainee struct into a map of DynamoDB attribute values
	item, err := attributevalue.MarshalMap(trainee)
	if err != nil {
		return fmt.Errorf("failed to marshal trainee: %w", err)
	}

	// Create the PutItem input
	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	}

	// Execute the PutItem operation
	_, err = r.client.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to save trainee to DynamoDB: %w", err)
	}

	return nil
}

// Update updates an existing trainee record in the DynamoDB table and returns an error if the operation fails.
func (r *TraineeDDBRepository) Update(trainee *models.Trainee) error {
	ctx := context.Background()

	// Check if the trainee exists before updating
	existingTrainee, err := r.FindByID(trainee.ID)
	if err != nil {
		return fmt.Errorf("trainee not found for update: %w", err)
	}

	if existingTrainee == nil {
		return fmt.Errorf("trainee with id %s does not exist", trainee.ID)
	}

	// Marshal the trainee struct into a map of DynamoDB attribute values
	item, err := attributevalue.MarshalMap(trainee)
	if err != nil {
		return fmt.Errorf("failed to marshal trainee for update: %w", err)
	}

	// Create the PutItem input (PutItem for a full update of the item)
	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	}

	// Execute the PutItem operation
	_, err = r.client.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update trainee in DynamoDB: %w", err)
	}

	return nil
}

// Delete removes a trainee record from the DynamoDB table based on the provided ID and returns an error if the operation fails.
func (r *TraineeDDBRepository) Delete(id string) error {
	ctx := context.Background()

	// Create the DeleteItem input
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	}

	// Execute the DeleteItem operation
	_, err := r.client.DeleteItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete trainee from DynamoDB: %w", err)
	}

	return nil
}

// DeleteByEmail removes a trainee record from the DynamoDB table based on the provided email and returns an error if the operation fails.
func (r *TraineeDDBRepository) DeleteByEmail(email string) error {
	ctx := context.Background()

	// Find the trainee by email first to get the ID
	trainee, err := r.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to find trainee for deletion: %w", err)
	}

	// Create the DeleteItem input using the ID
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: trainee.ID},
		},
	}

	// Execute the DeleteItem operation
	_, err = r.client.DeleteItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete trainee from DynamoDB: %w", err)
	}

	return nil
}

// CompleteTraining updates a trainee's record to mark training as complete.
func (r *TraineeDDBRepository) CompleteTraining(id string) error {
	ctx := context.Background()

	// Get the current timestamp for last training field
	currentTime := time.Now().Format(time.RFC3339)

	// Create the UpdateItem input
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET last_training = :lastTraining"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":lastTraining": &types.AttributeValueMemberS{Value: currentTime},
		},
	}

	// Execute the UpdateItem operation
	_, err := r.client.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update trainee training completion in DynamoDB: %w", err)
	}

	return nil
}

// Checkin updates a trainee's record to mark them as checked in.
func (r *TraineeDDBRepository) Checkin(id string) error {
	ctx := context.Background()

	// Create the UpdateItem input
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET checked_in = :checkedIn"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":checkedIn": &types.AttributeValueMemberBOOL{Value: true},
		},
	}

	// Execute the UpdateItem operation
	_, err := r.client.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to checkin trainee in DynamoDB: %w", err)
	}

	return nil
}

// Checkout updates a trainee's record to mark them as checked out.
func (r *TraineeDDBRepository) Checkout(id string) error {
	ctx := context.Background()

	// Create the UpdateItem input
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET checked_in = :checkedIn"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":checkedIn": &types.AttributeValueMemberBOOL{Value: false},
		},
	}

	// Execute the UpdateItem operation
	_, err := r.client.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to checkout trainee in DynamoDB: %w", err)
	}

	return nil
}

// FindAll retrieves all trainee records from the DynamoDB table and returns a slice of Trainee objects or an error.
func (r *TraineeDDBRepository) FindAll() ([]*models.Trainee, error) {
	ctx := context.Background()

	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	var trainees []*models.Trainee
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &trainees); err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainees: %w", err)
	}

	return trainees, nil
}
