package api

import (
	"log"
	"net/http"
	"photon-server/pkg/model"

	"github.com/aws/aws-sdk-go/aws"
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
	result, err := file.UploadForS3()

	if err == nil {
		body := model.FileUploadForS3{URL: result.Location}
		c.JSON(http.StatusOK, body)
		log.Println("file uploaded to " + aws.StringValue(&result.Location))
	} else {
		body := model.FileUploadForS3{URL: ""}
		c.JSON(http.StatusBadRequest, body)
	}
}
