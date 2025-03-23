package toolkit

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const randomStringSource = "abcdefghijklmnopqrstubwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type will have access
// to all the methods with the receiver pointer to tools
type Tools struct {
	MaxFileSize      int
	AllowedFileTypes []string
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

func (t *Tools) UploadedOneFile(r *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}
	files, err := t.UploadedFiles(r, uploadDir, renameFile)
	if err != nil {
		return nil, err
	}
	return files[0], nil
}

func (t *Tools) UploadedFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
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

//
//// DynamoDBAPI is an interface defining methods for interacting with Amazon DynamoDB.
//// It includes operations for retrieving, scanning, inserting, deleting, updating, and querying items.
//type DynamoDBAPI interface {
//	GetItem(ctx context.Context, params *ddb.GetItemInput, optFns ...func(*ddb.Options)) (*ddb.GetItemOutput, error)
//	Scan(ctx context.Context, params *ddb.ScanInput, optFns ...func(*ddb.Options)) (*ddb.ScanOutput, error)
//	PutItem(ctx context.Context, params *ddb.PutItemInput, optFns ...func(*ddb.Options)) (*ddb.PutItemOutput, error)
//	DeleteItem(ctx context.Context, params *ddb.DeleteItemInput, optFns ...func(*ddb.Options)) (*ddb.DeleteItemOutput, error)
//	UpdateItem(ctx context.Context, params *ddb.UpdateItemInput, optFns ...func(*ddb.Options)) (*ddb.UpdateItemOutput, error)
//	Query(ctx context.Context, params *ddb.QueryInput, optFns ...func(*ddb.Options)) (*ddb.QueryOutput, error)
//}
//
//// InitLambda initializes the AWS Lambda environment by setting up the DynamoDB client and required configurations.
//func (t *Tools) InitLambda(initialized bool, tableName string, ddbClient DynamoDBAPI) error {
//	if initialized {
//		return nil
//	}
//
//	// Get table name from environment
//	tableName = os.Getenv("DYNAMODB_TABLE_NAME")
//	if tableName == "" {
//		return fmt.Errorf("DYNAMODB_TABLE_NAME environment variable not set")
//	}
//
//	// Load AWS configuration and create a DynamoDB client
//	cfg, err := config.LoadDefaultConfig(context.TODO())
//	if err != nil {
//		return fmt.Errorf("unable to load SDK config: %w", err)
//	}
//	ddbClient = ddb.NewFromConfig(cfg)
//
//	initialized = true
//	return nil
//}
