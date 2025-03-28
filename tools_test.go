package toolkit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

// RoundTripFunc is a type that represents a function handling HTTP requests and returning HTTP responses.
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip executes the RoundTripFunc to handle the given HTTP request and returns the corresponding HTTP response and error.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient creates a new instance of *http.Client using the provided RoundTripFunc as the transport layer.
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// TestTools_PushJSONToRemote tests the PushJSONToRemote function by sending JSON data to a remote URL using a mock client.
// It verifies the response status code and ensures no errors occur during the process.
func TestTools_PushJSONToRemote(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		//	test request parameters
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"foo": "bar"}`)),
			Header:     make(http.Header),
		}
	})
	var testTools Tools
	var foo struct {
		Bar string `json:"bar"`
	}
	foo.Bar = "bar"

	_, _, err := testTools.PushJSONToRemote("http://example.com/some/path", foo, client)
	if err != nil {
		t.Error("Failed to call remote url", err)
	}

}

// TestTools_RandomString tests the RandomString function by verifying its output length for a given input value.
func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("String length is not 10")
	}
	t.Log(s)
}

// uploadTests is a test table for file upload scenarios, specifying test name, allowed file types, renaming behavior, and error expectation.
var uploadTests = []struct {
	name          string
	allowedTypes  []string
	renameFile    bool
	errorExpected bool
}{
	{name: "allowed no rename", allowedTypes: []string{"image/jpeg", "image/png", "image/gif"}, renameFile: false, errorExpected: false},
	{name: "allowed rename", allowedTypes: []string{"image/jpeg", "image/png", "image/gif"}, renameFile: true, errorExpected: false},
	{name: "not allowed", allowedTypes: []string{"image/jpeg", "image/gif"}, renameFile: false, errorExpected: true},
}

// TestTools_UploadFiles tests the UploadFiles method by simulating different upload scenarios and validating expected outcomes.
func TestTools_UploadFiles(t *testing.T) {
	for _, e := range uploadTests {
		// set up a pipe to avoid buffering
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer writer.Close()
			defer wg.Done()

			//create the for data field
			part, err := writer.CreateFormFile("file", "./testdata/image.png")
			if err != nil {
				t.Error(err)
			}

			f, err := os.Open("./testdata/image.png")
			if err != nil {
				t.Error(err)
			}
			defer f.Close()

			img, _, err := image.Decode(f)
			if err != nil {
				t.Error("Error decoding image", err)
			}
			err = png.Encode(part, img)
			if err != nil {
				t.Error(err)
			}
		}()

		// read from the pipe which receives data
		request := httptest.NewRequest("POST", "/", pr)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		var testTools Tools
		testTools.AllowedFileTypes = e.allowedTypes

		uploadedFiles, err := testTools.UploadFiles(request, "./testdata/uploads/", e.renameFile)
		if err != nil && !e.errorExpected {
			t.Error(err)
		}

		if !e.errorExpected {
			if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName)); os.IsNotExist(err) {
				t.Errorf("%s expected file to exist %s", e.name, err.Error())
			}

			// clean up
			_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName))
		}

		if !e.errorExpected && err != nil {
			t.Errorf("%s: error expected but none recieved", e.name)
		}

		wg.Wait()
	}
}

// TestTools_UploadOneFile is a test function to verify the behavior of the UploadOneFile method in various scenarios.
func TestTools_UploadOneFile(t *testing.T) {
	for _, e := range uploadTests {
		// set up a pipe to avoid buffering
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)

		go func() {
			defer writer.Close()

			//create the for data field
			part, err := writer.CreateFormFile("file", "./testdata/image.png")
			if err != nil {
				t.Error(err)
			}

			f, err := os.Open("./testdata/image.png")
			if err != nil {
				t.Error(err)
			}
			defer f.Close()

			img, _, err := image.Decode(f)
			if err != nil {
				t.Error("Error decoding image", err)
			}
			err = png.Encode(part, img)
			if err != nil {
				t.Error(err)
			}
		}()

		// read from the pipe which receives data
		request := httptest.NewRequest("POST", "/", pr)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		var testTools Tools

		uploadedFile, err := testTools.UploadOneFile(request, "./testdata/uploads/", true)
		if err != nil {
			t.Error(err)
		}

		if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFile.NewFileName)); os.IsNotExist(err) {
			t.Errorf("%s expected file to exist %s", e.name, err.Error())
		}

		// clean up
		_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFile.NewFileName))
	}
}

// TestTools_CreateDirIfNotExist tests the CreateDirIfNotExist method by creating and removing a directory to ensure proper functionality.
func TestTools_CreateDirIfNotExist(t *testing.T) {
	var testTools Tools
	err := testTools.CreateDirIfNotExist("./testdata/uploads")
	if err != nil {
		t.Error(err)
	}

	err = os.Remove("./testdata/uploads")
	if err != nil {
		t.Error(err)
	}
}

var slugTests = []struct {
	name          string
	s             string
	expected      string
	errorExpected bool
}{
	{"valid string", "Hello----------World!!!!! friend", "hello-world-friend", false},
	{"empty string", "", "", true},
	{"complex string", "Now is the time for all GOOD men! + fish & such &^123", "now-is-the-time-for-all-good-men-fish-such-123", false},
	{"japanese string", "こんにちは世界", "", true},
	{"japanese string and roman characters", "こんにちは世界 hello world", "hello-world", false},
}

// TestTools_Slugify tests the Slugify method of the Tools type to ensure it converts strings to valid URL-friendly slugs.
// It validates the correctness of slug generation against predefined test cases and checks for expected errors.
func TestTools_Slugify(t *testing.T) {
	var testTool Tools
	for _, e := range slugTests {
		slug, err := testTool.Slugify(e.s)

		if err != nil && !e.errorExpected {
			t.Errorf(`%s: error recieved when non expected: %s`, e.name, err.Error())
		}

		if !e.errorExpected && slug != e.expected {
			t.Errorf(`%s: wrong slug returned. expected: %s but recieved %s`, e.name, e.expected, slug)
		}
	}
}

// TestTools_DownloadStaticFile tests the DownloadStaticFile method by simulating an HTTP request and checking the response.
func TestTools_DownloadStaticFile(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	var testTools Tools

	testTools.DownloadStaticFile(rr, req, "./testdata", "pic.jpg", "puppy.jpg")
	res := rr.Result()
	defer res.Body.Close()

	if res.Header["Content-Length"][0] != "98827" {
		t.Error("Wrong content length", res.Header["Content-Length"][0])
	}

	if res.Header["Content-Disposition"][0] != "attachment; filename=\"puppy.jpg\"" {
		t.Error("Wrong content disposition", res.Header["Content-Disposition"][0])
	}

	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
}

// jsonTests is a test dataset for validating JSON parsing behavior with various scenarios including errors and constraints.
var jsonTests = []struct {
	name          string
	json          string
	errorExpected bool
	maxSize       int
	allowUnknown  bool
}{
	{"good json", `{"foo": "bar"}`, false, 1024, false},
	{"bad json", `{"foo": "bar"`, true, 1024, false},
	{"incorrect type", `{"foo": 1}`, true, 1024, false},
	{"two json files", `{"foo": "bar"}{"alpha": "beta"}`, true, 1024, false},
	{"empty body", "", true, 1024, false},
	{"syntax error in json", `{"foo": 1"`, true, 1024, false},
	{"unknown field in json", `{"fooo": "1"}`, true, 1024, false},
	{"allow unknown field in json", `{"fooo": "1"}`, false, 1024, true},
	{"too large json", `{"foo": "bar"}`, true, 1, false},
	{"missing field name", `{jack: "1"}`, true, 1024, true},
	{"not json", "hello world", true, 1024, true},
}

// TestTools_ReadJSON verifies the behavior of the Tools.ReadJSON method with various JSON inputs, sizes, and configurations.
func TestTools_ReadJSON(t *testing.T) {
	var testTool Tools

	for _, e := range jsonTests {
		//	set the max file size
		testTool.MaxJSONSize = e.maxSize

		// allow or disallow unknown fields
		testTool.AllowUnknownFields = e.allowUnknown

		//	create a variable to read the decoded json
		var decodedJSON struct {
			Foo string `json:"foo"`
		}

		// create a request
		req, err := http.NewRequest("POST", "/", bytes.NewReader([]byte(e.json)))
		if err != nil {
			t.Log("error:", err)
		}

		// create a recorder
		rr := httptest.NewRecorder()
		err = testTool.ReadJSON(rr, req, &decodedJSON)

		if e.errorExpected && err == nil {
			t.Errorf("%s: error expected but none recieved", e.name)
		}
		if !e.errorExpected && err != nil {
			t.Errorf("%s: error recieved when non expected: %s", e.name, err.Error())
		}

		// prevent resource leak
		err = req.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}
}

// TestTools_WriteJSON tests the WriteJSON function to ensure it correctly writes a JSON response with headers and status code.
func TestTools_WriteJSON(t *testing.T) {
	var testTool Tools
	rr := httptest.NewRecorder()
	payload := JSONResponse{
		Error:   false,
		Message: "foo",
	}

	headers := make(http.Header)
	headers.Add("FOO", "BAR")

	err := testTool.WriteJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("failed to write json: %v", err)
	}
}

// TestTools_ErrorJSON tests the ErrorJSON method to ensure it returns the correct JSON error response and status code.
func TestTools_ErrorJSON(t *testing.T) {
	var testTool Tools
	rr := httptest.NewRecorder()
	err := testTool.ErrorJSON(rr, errors.New("some error"), http.StatusServiceUnavailable)
	if err != nil {
		t.Error(err)
	}

	var payload JSONResponse
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		t.Error("received error while decoding json", err)
	}

	if !payload.Error {
		t.Error("error flag set to false but must be true")
	}

	if rr.Code != http.StatusServiceUnavailable {
		t.Errorf("wrong status code recieved: %d", rr.Code)
	}
}

// Test environment variables
func TestEnvironmentVariables(t *testing.T) {
	setupTest(t)

	// Test environment variable is set
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	assert.Equal(t, "test-table", tableName)

	// Test the table name is accessible
	assert.Equal(t, "test-table", safeGetTableName())
}

// Look at your main.go for a function like initLambda or initialize
// and test it like this (replace initLambda with your actual function name)
func TestInitDDBFunction(t *testing.T) {

	testDatabase := &Database{
		DBName:      "DYNAMODB_TABLE_NAME",
		Initialized: initialized,
		Table:       tableName,
	}
	// Save original state
	originalTableName := safeGetTableName()
	originalInit := safeGetInitialized()
	originalClient := safeGetDynamoDBClient()
	defer func() {
		// Restore original state after test
		safeSetTableName(originalTableName)
		safeSetInitialized(originalInit)
		safeSetDynamoDBClient(originalClient)
	}()

	// Reset state
	safeSetTableName("")
	safeSetInitialized(false)
	safeSetDynamoDBClient(nil)

	// Set environment variable
	err := os.Setenv("DYNAMODB_TABLE_NAME", "test-init-lambda")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}

	// Call your initialization function
	// Replace initLambda with your actual function name
	err = testDatabase.InitDDBLambda(ddbClient)
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	assert.NoError(t, err)

	// Reset environment variable
	err = os.Setenv("DYNAMODB_TABLE_NAME", "test-table")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
}
