package model

import (
	"encoding/base64"
	"log"
	"os"
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

// Create create file
func (f File) Create() error {
	data, _ := base64.StdEncoding.DecodeString(f.Base64)

	file, err := os.Create("./upload_files/" + f.Name)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer file.Close()

	file.Write(data)

	return err
}

// Delete delete file
func (f File) Delete() error {

	err := os.Remove("./upload_files/" + f.Name)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
