package model

import (
	"bytes"
	"encoding/base64"

	"photon-server/pkg/awsclient"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// File is struct
type File struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ContentType string `json:"contentType"`
	FileLength  uint64 `json:"fileLength"`
	Base64      string `json:"base64"`
}

// FileUploadForS3 is struct
type FileUploadForS3 struct {
	URL string `json:"url"`
}

// UploadForS3 uploadForS3 file
func (f File) UploadForS3() (*s3manager.UploadOutput, error) {
	data, _ := base64.StdEncoding.DecodeString(f.Base64)

	// Upload the file to S3.
	sess, err := awsclient.CreateAWSSession()
	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("photon-s3"),
		Key:    aws.String(f.ID + "/" + time.Now().Format("20060102150405") + ".jpg"),
		Body:   bytes.NewReader(data),
		ACL:    aws.String("public-read"),
	})

	return result, err
}
