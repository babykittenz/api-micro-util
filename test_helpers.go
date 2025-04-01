package toolkit

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
)

// Mutex for safe variable access during tests
var testMutex sync.RWMutex
var tableName string
var ddbClient DynamoDBAPI
var initialized bool = false

// safeSetTableName sets the tableName with mutex protection
func safeSetTableName(name string) {
	testMutex.Lock()
	defer testMutex.Unlock()

	log.Printf("Setting tableName to %q", name)
	tableName = name
}

// SafeGetTableName gets the tableName with mutex protection
func SafeGetTableName() string {
	// Try to get the environment-specific table name
	env := os.Getenv("ENVIRONMENT") // This could be "dev", "staging", "prod"
	if env == "" {
		env = "dev" // Default to dev if not specified
	}

	// Get the base table name
	baseTableName := os.Getenv("DYNAMODB_TABLE_NAME")
	if baseTableName == "" {
		baseTableName = "trainees" // Fallback to a default
	}

	// For dev and staging, append the environment to avoid collision
	// For prod, use the clean name (or you could still append if preferred)
	tableName := baseTableName
	if env != "prod" {
		tableName = fmt.Sprintf("%s_%s", baseTableName, env)
	}

	log.Printf("Using DynamoDB table: %s for environment: %s", tableName, env)
	return tableName
}

// safeSetDynamoDBClient sets the DynamoDB client with mutex protection
func safeSetDynamoDBClient(client DynamoDBAPI) {
	testMutex.Lock()
	defer testMutex.Unlock()

	ddbClient = client
}

// SafeGetDynamoDBClient gets the DynamoDB client with mutex protection
func SafeGetDynamoDBClient() DynamoDBAPI {
	testMutex.RLock()
	defer testMutex.RUnlock()

	return ddbClient
}

// safeSetInitialized sets the initialized flag with mutex protection
func safeSetInitialized(value bool) {
	testMutex.Lock()
	defer testMutex.Unlock()

	initialized = value
}

// safeGetInitialized gets the initialized flag with mutex protection
func safeGetInitialized() bool {
	testMutex.RLock()
	defer testMutex.RUnlock()

	return initialized
}

// ensureEnvironment ensures environment variables and package variables are set correctly
func ensureEnvironment(t *testing.T) {
	// Set environment variable
	err := os.Setenv("DYNAMODB_TABLE_NAME", "test-table")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}

	// Set package variables
	safeSetTableName("test-table")
	safeSetInitialized(true)

	// Only set ddbClient if it's nil
	if SafeGetDynamoDBClient() == nil {
		safeSetDynamoDBClient(&MockDynamoDBClient{t: t})
	}

	t.Logf("Environment ensured: tableName=%q, env=%q, initialized=%v",
		SafeGetTableName(),
		os.Getenv("DYNAMODB_TABLE_NAME"),
		safeGetInitialized())
}
