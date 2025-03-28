package toolkit

import (
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

// safeGetTableName gets the tableName with mutex protection
func safeGetTableName() string {
	testMutex.RLock()
	defer testMutex.RUnlock()

	return tableName
}

// safeSetDynamoDBClient sets the DynamoDB client with mutex protection
func safeSetDynamoDBClient(client DynamoDBAPI) {
	testMutex.Lock()
	defer testMutex.Unlock()

	ddbClient = client
}

// safeGetDynamoDBClient gets the DynamoDB client with mutex protection
func safeGetDynamoDBClient() DynamoDBAPI {
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
	if safeGetDynamoDBClient() == nil {
		safeSetDynamoDBClient(&MockDynamoDBClient{t: t})
	}

	t.Logf("Environment ensured: tableName=%q, env=%q, initialized=%v",
		safeGetTableName(),
		os.Getenv("DYNAMODB_TABLE_NAME"),
		safeGetInitialized())
}
