package toolkit

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const randomStringSource = "abcdefghijklmnopqrstubwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type will have access
// to all the methods with the receiver pointer to tools
type Tools struct {
	MaxFileSize        int
	AllowedFileTypes   []string
	MaxJSONSize        int
	AllowUnknownFields bool
}

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string
func (t *Tools) RandomString(length int) string {
	s, r := make([]rune, length), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

// UploadedFile is a struct used to save information about an uploaded file
type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int64
}

// UploadOneFile uploads a single file from an HTTP request, saves it to the specified directory, and returns its information.
func (t *Tools) UploadOneFile(r *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}
	files, err := t.UploadFiles(r, uploadDir, renameFile)
	if err != nil {
		return nil, err
	}
	return files[0], nil
}

// UploadFiles processes and saves uploaded files from an HTTP request to the specified directory.
// It optionally renames files with a random filename if `rename` is set to true.
// Returns a slice of uploaded file information or an error if the operation fails.
func (t *Tools) UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}
	var uploadedFiles []*UploadedFile

	if t.MaxFileSize == 0 {
		t.MaxFileSize = 1024 * 1024 * 1024 // about a gig
	}

	err := t.CreateDirIfNotExist(uploadDir)
	if err != nil {
		return nil, err
	}

	err = r.ParseMultipartForm(int64(t.MaxFileSize))
	if err != nil {
		return nil, errors.New("the uploaded file is too large")
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, hdr := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
				var uploadedFile UploadedFile
				// open the file
				infile, err := hdr.Open()
				// check for error
				if err != nil {
					return nil, err
				}
				// we have an infile we need to close to prevent resource leak
				defer func(infile multipart.File) {
					err := infile.Close()
					if err != nil {

					}
				}(infile)

				// check the first 512 bytes
				buff := make([]byte, 512)
				_, err = infile.Read(buff)
				if err != nil {
					return nil, err
				}

				// TODO: check to see if the file type is permitted
				allowed := false
				fileType := http.DetectContentType(buff)

				if len(t.AllowedFileTypes) > 0 {
					for _, allowedType := range t.AllowedFileTypes {
						if strings.EqualFold(fileType, allowedType) {
							allowed = true
							break
						}
					}
				} else {
					// in the hands of the user to not allow all file types
					allowed = true
				}

				if !allowed {
					return nil, errors.New("file type not allowed")
				}

				_, err = infile.Seek(0, 0)
				if err != nil {
					return nil, err
				}

				if renameFile {
					uploadedFile.NewFileName = fmt.Sprintf("%s%s", t.RandomString(25), filepath.Ext(hdr.Filename))
				} else {
					uploadedFile.NewFileName = hdr.Filename
				}

				uploadedFile.OriginalFileName = hdr.Filename
				var outfile *os.File
				defer func(outfile *os.File) {
					err := outfile.Close()
					if err != nil {

					}
				}(outfile)

				if outfile, err = os.Create(filepath.Join(uploadDir, uploadedFile.NewFileName)); err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outfile, infile)
					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}
				uploadedFiles = append(uploadedFiles, &uploadedFile)
				return uploadedFiles, nil
			}(uploadedFiles)
			if err != nil {
				return uploadedFiles, err
			}
		}
	}
	return uploadedFiles, nil
}

// CreateDirIfNotExist creates a directory at the specified path if it does not already exist, using permission mode 0755.
// Returns an error if directory creation fails.
func (t *Tools) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// Slugify converts a string into a URL-friendly slug by replacing non-alphanumeric characters with hyphens and trimming excess hyphens.
// Returns the slugified string or an error if the input results in an empty or invalid slug.
func (t *Tools) Slugify(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string not allowed")
	}

	var re = regexp.MustCompile(`[^a-z\d]+`)

	slug := strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
	if len(slug) == 0 {
		return "", errors.New("after removing characters slug is zero length")
	}

	return slug, nil
}

// DownloadStaticFile downloads a file, and tries to force the browser to avoid displaying it
// in the browser window by setting content disposition. It also allows specification of the
// display name
func (t *Tools) DownloadStaticFile(w http.ResponseWriter, r *http.Request, p, file, displayName string) {
	fp := path.Join(p, file)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", displayName))

	http.ServeFile(w, r, fp)
}

// JSONResponse represents a standard JSON response structure with error, message, and optional data fields.
type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ReadJSON reads JSON data from an HTTP request body into the provided data structure.
// It ensures the body size does not exceed MaxJSONSize and validates for allowed fields.
// Returns an error if the JSON format is invalid, too large, or contains unknown fields.
func (t *Tools) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // about 1 meg
	if t.MaxJSONSize != 0 {
		maxBytes = t.MaxJSONSize
	}
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	if !t.AllowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(data)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Value != "" {
				return fmt.Errorf("body contains an incorrect JSON type for field %q ", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains an incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must contain only one JSON object")
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		case errors.As(err, &invalidUnmarshalError):
			return fmt.Errorf("body contains an incorrect JSON type for field %q ", invalidUnmarshalError.Type.String())
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain only one JSON object")
	}

	return nil
}

// WriteJSON writes the given data as JSON to the HTTP response with the specified status code and optional headers.
func (t *Tools) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// ErrorJSON sends a JSON error response with the specified status code or a default status of Bad Gateway (502).
func (t *Tools) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadGateway
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return t.WriteJSON(w, statusCode, payload)
}

// PushJSONToRemote posts arbitrary data to some URL as JSON, and returns the response, status code, and error, if any.
// The final parameter, client, is optional. If none is specified, we use the standard http.Client.
func (t *Tools) PushJSONToRemote(uri string, data interface{}, client ...*http.Client) (*http.Response, int, error) {
	// create json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, 0, err
	}

	// check for custom http client
	httpClient := &http.Client{}
	if len(client) > 0 {
		httpClient = client[0]
	}

	// build the request and set the header
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, 0, err
	}
	request.Header.Set("Content-Type", "application/json")

	// call the remote uri
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	// send response back
	return response, response.StatusCode, nil
}

// Database TODO: need to create a repository pattern / adapter pattern
type Database struct {
	Initialized bool
	DBName      string
	Table       string
}

// DynamoDBAPI is an interface defining methods for interacting with Amazon DynamoDB.
// It includes operations for retrieving, scanning, inserting, deleting, updating, and querying items.
type DynamoDBAPI interface {
	GetItem(ctx context.Context, params *ddb.GetItemInput, optFns ...func(*ddb.Options)) (*ddb.GetItemOutput, error)
	Scan(ctx context.Context, params *ddb.ScanInput, optFns ...func(*ddb.Options)) (*ddb.ScanOutput, error)
	PutItem(ctx context.Context, params *ddb.PutItemInput, optFns ...func(*ddb.Options)) (*ddb.PutItemOutput, error)
	DeleteItem(ctx context.Context, params *ddb.DeleteItemInput, optFns ...func(*ddb.Options)) (*ddb.DeleteItemOutput, error)
	UpdateItem(ctx context.Context, params *ddb.UpdateItemInput, optFns ...func(*ddb.Options)) (*ddb.UpdateItemOutput, error)
	Query(ctx context.Context, params *ddb.QueryInput, optFns ...func(*ddb.Options)) (*ddb.QueryOutput, error)
}

// InitDDBLambda initializes the DynamoDB client and retrieves the table name from the environment if not already initialized.
// It ensures the initialization is performed only once to avoid redundant operations.
// Returns an error if the environment variable for the table name is not set or the AWS SDK configuration fails to load.
func (d *Database) InitDDBLambda(ddbClient DynamoDBAPI) error {
	if d.Initialized {
		return nil
	}

	// Get table name from environment
	d.Table = os.Getenv(d.DBName)
	if d.Table == "" {
		return fmt.Errorf("%s environment variable not set", d.DBName)
	}

	// Load AWS configuration and create a DynamoDB client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %w", err)
	}
	ddbClient = ddb.NewFromConfig(cfg)

	d.Initialized = true
	return nil
}
