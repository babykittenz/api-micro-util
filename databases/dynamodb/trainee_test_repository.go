package dynamodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/babykittenz/api-micro-util"
	"github.com/babykittenz/api-micro-util/models"
	"github.com/babykittenz/api-micro-util/repository"
)

type testTraineeDDBRepository struct {
	client    *toolkit.MockDynamoDBClient
	tableName string
	tools     *toolkit.Tools
}

func NewTestTraineeDDBRepository(client *toolkit.MockDynamoDBClient) repository.TraineeRepository {
	return &testTraineeDDBRepository{
		client:    client,
		tableName: toolkit.SafeGetTableName(),
		tools:     &toolkit.Tools{},
	}
}

// FindByID retrieves a trainee record from DynamoDB based on the provided ID
func (r *testTraineeDDBRepository) FindByID(id string) (*models.Trainee, error) {
	// TODO: move context to test setup not called every function
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
	if result.Item == nil || len(result.Item) == 0 {
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

// FindByEmail retrieves a trainee record from DynamoDB based on the provided email
func (r *testTraineeDDBRepository) FindByEmail(email string) (*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for email
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainee from DynamoDB: %w", err)
	}

	// If no item found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with email %s not found", email)
	}

	// Unmarshal the first matching item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByPhone retrieves a trainee record from DynamoDB based on the provided phone number
func (r *testTraineeDDBRepository) FindByPhone(phone string) (*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for phone
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("phone = :phone"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":phone": &types.AttributeValueMemberS{Value: phone},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainee from DynamoDB: %w", err)
	}

	// If no item found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with phone %s not found", phone)
	}

	// Unmarshal the first matching item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByPhoneAndLocation retrieves a trainee record based on phone and location ID
func (r *testTraineeDDBRepository) FindByPhoneAndLocation(phone string, locationID string) (*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for phone and location_id
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("phone = :phone AND location_id = :location_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":phone":       &types.AttributeValueMemberS{Value: phone},
			":location_id": &types.AttributeValueMemberS{Value: locationID},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainee from DynamoDB: %w", err)
	}

	// If no item found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with phone %s and location %s not found", phone, locationID)
	}

	// Unmarshal the first matching item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByEmailAndLocation retrieves a trainee record based on email and location ID
func (r *testTraineeDDBRepository) FindByEmailAndLocation(email string, locationID string) (*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for email and location_id
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("email = :email AND location_id = :location_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email":       &types.AttributeValueMemberS{Value: email},
			":location_id": &types.AttributeValueMemberS{Value: locationID},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainee from DynamoDB: %w", err)
	}

	// If no item found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with email %s and location %s not found", email, locationID)
	}

	// Unmarshal the first matching item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByNames retrieves a trainee record based on first and last names
func (r *testTraineeDDBRepository) FindByNames(firstName string, lastName string) (*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for first_name and last_name
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("first_name = :first_name AND last_name = :last_name"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":first_name": &types.AttributeValueMemberS{Value: firstName},
			":last_name":  &types.AttributeValueMemberS{Value: lastName},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainee from DynamoDB: %w", err)
	}

	// If no item found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with name %s %s not found", firstName, lastName)
	}

	// Unmarshal the first matching item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindByNamesAndLocation retrieves a trainee record based on names and location ID
func (r *testTraineeDDBRepository) FindByNamesAndLocation(firstName string, lastName string, locationID string) (*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for first_name, last_name and location_id
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("first_name = :first_name AND last_name = :last_name AND location_id = :location_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":first_name":  &types.AttributeValueMemberS{Value: firstName},
			":last_name":   &types.AttributeValueMemberS{Value: lastName},
			":location_id": &types.AttributeValueMemberS{Value: locationID},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainee from DynamoDB: %w", err)
	}

	// If no item found
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("trainee with name %s %s and location %s not found", firstName, lastName, locationID)
	}

	// Unmarshal the first matching item
	var trainee models.Trainee
	err = attributevalue.UnmarshalMap(result.Items[0], &trainee)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
	}

	return &trainee, nil
}

// FindAll retrieves all trainee records from the DynamoDB table
func (r *testTraineeDDBRepository) FindAll() ([]*models.Trainee, error) {
	ctx := context.Background()

	// Scan the entire table
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Convert the DynamoDB items to Trainee structs
	trainees := make([]*models.Trainee, 0, len(result.Items))
	for _, item := range result.Items {
		var trainee models.Trainee
		err = attributevalue.UnmarshalMap(item, &trainee)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
		}
		trainees = append(trainees, &trainee)
	}

	return trainees, nil
}

// FindAllByCompanyID retrieves all trainee records associated with a company ID
func (r *testTraineeDDBRepository) FindAllByCompanyID(id string) ([]*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for company_id
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("company_id = :company_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":company_id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Convert the DynamoDB items to Trainee structs
	trainees := make([]*models.Trainee, 0, len(result.Items))
	for _, item := range result.Items {
		var trainee models.Trainee
		err = attributevalue.UnmarshalMap(item, &trainee)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
		}
		trainees = append(trainees, &trainee)
	}

	return trainees, nil
}

// FindAllByRegionID retrieves all trainee records associated with a region ID
func (r *testTraineeDDBRepository) FindAllByRegionID(id string) ([]*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for region_id
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("region_id = :region_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":region_id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Convert the DynamoDB items to Trainee structs
	trainees := make([]*models.Trainee, 0, len(result.Items))
	for _, item := range result.Items {
		var trainee models.Trainee
		err = attributevalue.UnmarshalMap(item, &trainee)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
		}
		trainees = append(trainees, &trainee)
	}

	return trainees, nil
}

// FindAllByLocationID retrieves all trainee records associated with a location ID
func (r *testTraineeDDBRepository) FindAllByLocationID(id string) ([]*models.Trainee, error) {
	ctx := context.Background()

	// Set up scan with filter expression for location_id
	result, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("location_id = :location_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":location_id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan trainees from DynamoDB: %w", err)
	}

	// Convert the DynamoDB items to Trainee structs
	trainees := make([]*models.Trainee, 0, len(result.Items))
	for _, item := range result.Items {
		var trainee models.Trainee
		err = attributevalue.UnmarshalMap(item, &trainee)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal trainee: %w", err)
		}
		trainees = append(trainees, &trainee)
	}

	return trainees, nil
}

// Save persists a trainee record to the DynamoDB table
func (r *testTraineeDDBRepository) Save(trainee *models.Trainee) error {
	if trainee == nil {
		return errors.New("cannot save nil trainee")
	}

	ctx := context.Background()

	// Marshal the trainee struct to DynamoDB attribute values
	item, err := attributevalue.MarshalMap(trainee)
	if err != nil {
		return fmt.Errorf("failed to marshal trainee: %w", err)
	}

	// Save the item to DynamoDB
	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	})

	if err != nil {
		return fmt.Errorf("failed to save trainee to DynamoDB: %w", err)
	}

	return nil
}

// Update updates an existing trainee record in the DynamoDB table
func (r *testTraineeDDBRepository) Update(trainee *models.Trainee) error {
	if trainee == nil {
		return errors.New("cannot update nil trainee")
	}

	// Check if trainee exists
	_, err := r.FindByID(trainee.ID)
	if err != nil {
		return fmt.Errorf("trainee to update not found: %w", err)
	}

	// If exists, save the updated record
	return r.Save(trainee)
}

// Delete removes a trainee record from the DynamoDB table
func (r *testTraineeDDBRepository) Delete(id string) error {
	ctx := context.Background()

	// Delete the item from DynamoDB
	_, err := r.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return fmt.Errorf("failed to delete trainee from DynamoDB: %w", err)
	}

	return nil
}

// DeleteByEmail removes a trainee record from the DynamoDB table by email
func (r *testTraineeDDBRepository) DeleteByEmail(email string) error {
	// First find the trainee by email
	trainee, err := r.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to find trainee to delete: %w", err)
	}

	// Then delete by ID
	return r.Delete(trainee.ID)
}

// CompleteTraining marks the training session as complete for a trainee
func (r *testTraineeDDBRepository) CompleteTraining(id string) error {
	// Find the trainee first
	trainee, err := r.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find trainee to complete training: %w", err)
	}

	// Set the last training timestamp to current time
	trainee.LastTraining = fmt.Sprintf("%d", r.tools.GetCurrentTimestamp())

	// Save the updated trainee
	return r.Save(trainee)
}

// Checkin marks a trainee as checked in
func (r *testTraineeDDBRepository) Checkin(id string) error {
	// Find the trainee first
	trainee, err := r.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find trainee to checkin: %w", err)
	}

	// Set the checked in status
	trainee.CheckedIn = true

	// Save the updated trainee
	return r.Save(trainee)
}

// Checkout marks a trainee as checked out
func (r *testTraineeDDBRepository) Checkout(id string) error {
	// Find the trainee first
	trainee, err := r.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find trainee to checkout: %w", err)
	}

	// Set the checked in status to false
	trainee.CheckedIn = false

	// Save the updated trainee
	return r.Save(trainee)
}
