# API Micro Util

A simple example of how to create a reusable Go module with commonly used tools and repository implementations.

## Included Tools

- [X] Read JSON
- [X] Write JSON
- [X] Produce a JSON encoded error response
- [X] Upload a file to a specified directory
- [X] Download a static file
- [X] Get a random string of length n
- [X] Post JSON to a remote service
- [X] Create a directory, including all parent directories, if it does not already exist
- [X] Create a URL safe slug from a string

## Repository Features

- [X] DynamoDB repository implementations for various models:
    - Company
    - Region
    - Location
    - Trainee
    - User
    - Training
    - Checkin
    - TextMessage
    - AutomaticTextMessage
    - Language
- [X] Test implementations for all repositories
- [X] Comprehensive mock DynamoDB client for testing

## Testing Support

- [X] Mock DynamoDB client for unit testing
- [X] Test repository implementations
- [X] Test utilities and helpers
- [X] Thread-safe test environment management

## Installation

```bash
go get -u github.com/babykittenz/api-micro-util
```

## Example Usage

### Using the Tools

```go
package main

import (
    "net/http"
    "github.com/babykittenz/api-micro-util/toolkit"
)

func main() {
    tools := toolkit.Tools{}
    
    // Generate a random string
    randomString := tools.RandomString(10)
    
    // Create a directory
    err := tools.CreateDirIfNotExist("./data")
    if err != nil {
        // Handle error
    }
}
```

### Using the Repository Pattern

```go
package main

import (
    "github.com/babykittenz/api-micro-util/models"
    "github.com/babykittenz/api-micro-util/repository/dynamodb"
)

func main() {
    // Get a DynamoDB client (in production)
    client := GetDynamoDBClient()
    
    // Initialize repository
    traineeRepo := dynamodb.NewTraineeDDBRepository(client)
    
    // Find a trainee
    trainee, err := traineeRepo.FindByID("trainee-123")
    if err != nil {
        // Handle error
    }
    
    // Save a trainee
    trainee.FirstName = "John"
    err = traineeRepo.Save(trainee)
    if err != nil {
        // Handle error
    }
}
```

### Using Test Repositories

```go
package main_test

import (
    "testing"
    "github.com/babykittenz/api-micro-util/models"
    "github.com/babykittenz/api-micro-util/repository/dynamodb"
    "github.com/babykittenz/api-micro-util/toolkit"
)

func TestTraineeOperations(t *testing.T) {
    // Get a mock DynamoDB client
    mockClient := &toolkit.MockDynamoDBClient{T: t}
    
    // Initialize test repository
    repo := dynamodb.NewTestTraineeDDBRepository(mockClient)
    
    // Test repository operations
    trainee, err := repo.FindByID("test-id")
    
    // Make assertions
    // ...
}
```

## License

This project is licensed under the MIT License.