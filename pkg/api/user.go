package api

import (
	"Photon_Server/pkg/model"
	"Photon_Server/pkg/repository"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserAPI user api
type UserAPI struct {
	repository repository.UserRepository
}

// NewUserAPI user api create
func NewUserAPI(router *gin.RouterGroup, repo repository.UserRepository) {
	userAPI := &UserAPI{
		repository: repo,
	}
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userAPI.GetAll)
		userRoutes.GET("/:id", userAPI.GetByID)
		userRoutes.POST("", userAPI.Create)
		userRoutes.PUT("", userAPI.Update)
		userRoutes.DELETE("/:id", userAPI.Delete)
	}
}

// GetAll 複数のUserを取得します
func (api *UserAPI) GetAll(c *gin.Context) {
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

// GetByID UserをIDで検索して取得します
func (api *UserAPI) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := api.repository.GetByID(id)

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

// Create Userを作成します
func (api *UserAPI) Create(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	id, err := api.repository.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, model.CreatedResponse{ID: id})
}

// Update Userを更新します。
func (api *UserAPI) Update(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	err := api.repository.Update(user)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Delete Userを削除します
func (api *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := api.repository.Delete(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
