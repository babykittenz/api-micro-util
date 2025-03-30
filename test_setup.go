package toolkit

import (
	"log"
	"os"
	"testing"
	"time"
)

// init runs before any tests, including TestMain
func init() {
	log.Println("Initializing in init()...")

	// Set environment variable
	err := os.Setenv("DYNAMODB_TABLE_NAME", "test-table")
	if err != nil {
		log.Fatal(err)
	}
}

// TestMain handles common setup for all tests
func TestMain(m *testing.M) {
	log.Println("Initializing test environment in TestMain()...")

	// Set environment variables again to be sure
	err := os.Setenv("DYNAMODB_TABLE_NAME", "test-table")
	if err != nil {
		log.Fatal(err)
	}

	// Directly set the package variables using thread-safe accessors
	safeSetTableName("test-table")

	// Create the mock client
	safeSetDynamoDBClient(&MockDynamoDBClient{})

	// Mark as initialized
	safeSetInitialized(true)

	log.Printf("Test environment initialized: tableName=%s, ddbClient initialized=%v\n",
		SafeGetTableName(), safeGetDynamoDBClient() != nil)

	// Add a small delay to ensure variables propagate
	time.Sleep(50 * time.Millisecond)

	// Run the tests
	exitCode := m.Run()

	// Clean up
	os.Exit(exitCode)
}

// setupTest ensures test environment is properly set for each test
func setupTest(t *testing.T) {
	// Always set these variables for each test to ensure consistency
	t.Logf("Setting up test environment... (current tableName: %q)", SafeGetTableName())

	// Force-set table name for every test
	safeSetTableName("test-table")
	t.Logf("Set tableName to %q", SafeGetTableName())

	// Ensure environment variable is set
	if os.Getenv("DYNAMODB_TABLE_NAME") != "test-table" {
		err := os.Setenv("DYNAMODB_TABLE_NAME", "test-table")
		if err != nil {
			t.Fatalf("Failed to set environment variable: %v", err)
		}
		t.Logf("Reset environment variable DYNAMODB_TABLE_NAME to 'test-table'")
	}

	if !safeGetInitialized() {
		t.Log("Setting initialized=true")
		safeSetInitialized(true)
	}

	if safeGetDynamoDBClient() == nil {
		t.Log("Creating mock DynamoDB client")
		safeSetDynamoDBClient(&MockDynamoDBClient{t: t})
	}

	t.Logf("Test setup complete - tableName: %q, env: %q, initialized: %v",
		SafeGetTableName(),
		os.Getenv("DYNAMODB_TABLE_NAME"),
		safeGetInitialized())
}
