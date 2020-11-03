package api

import (
	"database/sql"
	"log"
	"net/http"
	"photon-server/pkg/repository"

	"github.com/gin-gonic/gin"
)

// PostAPI user api
type PostAPI struct {
	repository repository.PostRepository
}

// NewPostAPI user api create
func NewPostAPI(router *gin.RouterGroup, repo repository.PostRepository) {
	postAPI := &PostAPI{
		repository: repo,
	}
	userRoutes := router.Group("/post")
	{
		userRoutes.GET("", postAPI.GetAll)
	}
}

// GetAll 全ての写真投稿を取得します
func (api *PostAPI) GetAll(c *gin.Context) {
	result, err := api.repository.GetAll()

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
