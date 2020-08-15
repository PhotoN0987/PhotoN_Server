package api

import (
	"log"
	"net/http"
	"os"
	"photon-server/pkg/model"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

// FileAPI file api
type FileAPI struct{}

// NewFileAPI file api create
func NewFileAPI(router *gin.RouterGroup) {
	fileAPI := FileAPI{}
	fileUploadRoutes := router.Group("/file")
	{
		fileUploadRoutes.POST("", fileAPI.UploadFileForS3)
	}
}

// UploadFileForS3 ファイルデータをs3へUploadします。
func (api *FileAPI) UploadFileForS3(c *gin.Context) {

	var file model.File
	c.BindJSON(&file)
	err := file.Create()

	// file Open
	uploadFile, err := os.Open("./upload_files/" + file.Name)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "ファイルのアップロードに失敗しました")
		return
	}
	defer uploadFile.Close()

	// Upload the file to S3.
	sess, err := model.CreateAWSSession()
	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("photon-s3"),
		Key:    aws.String(file.ID + "/" + time.Now().Format("20060102150405") + ".jpg"),
		Body:   uploadFile,
		ACL:    aws.String("public-read"),
	})
	if err == nil {
		body := model.FileUploadForS3{URL: result.Location}
		c.JSON(http.StatusOK, body)
		log.Println("file uploaded to " + aws.StringValue(&result.Location))
	} else {
		body := model.FileUploadForS3{URL: ""}
		c.JSON(http.StatusBadRequest, body)
	}

	file.Delete()
}
