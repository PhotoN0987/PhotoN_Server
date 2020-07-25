package api_test

import (
	"Photon_Server/pkg/api"
	"Photon_Server/pkg/mock"
	"Photon_Server/pkg/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func newMockRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	apiV1 := router.Group("api/v1")

	return router, apiV1
}

func TestGetAll(t *testing.T) {
	mockUsers := make([]model.UserTable, 0)
	mockUser := model.UserTable{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()
	mockUsers = append(mockUsers, mockUser)

	gin.SetMode(gin.TestMode)

	mockUserRepo := new(mock.UserRepository)

	mockUserRepo.On("GetAll").Return(mockUsers, nil).Once()

	router, rg := newMockRouter()
	api.NewUserAPI(rg, mockUserRepo)

	getAllRes := httptest.NewRecorder()
	getAllReq, _ := http.NewRequest("GET", "/api/v1/users", nil)
	router.ServeHTTP(getAllRes, getAllReq)

	assert.Equal(t, 200, getAllRes.Code)
}

func TestGetByID(t *testing.T) {
	mockUser := model.UserTable{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()

	gin.SetMode(gin.TestMode)

	mockUserRepo := new(mock.UserRepository)

	mockUserRepo.On("GetByID", mockUser.ID).Return(mockUser, nil).Once()

	router, rg := newMockRouter()
	api.NewUserAPI(rg, mockUserRepo)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/"+strconv.Itoa(mockUser.ID), nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestCreate(t *testing.T) {
	mockUser := model.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	gin.SetMode(gin.TestMode)

	mockUserRepo := new(mock.UserRepository)
	mockUserRepo.On("Create", mockUser).Return(int64(mockUser.ID), nil).Once()

	router, rg := newMockRouter()
	api.NewUserAPI(rg, mockUserRepo)

	user_json, _ := json.Marshal(mockUser)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewReader(user_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 201, res.Code)
}

func TestUpdate(t *testing.T) {
	mockUser := model.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	gin.SetMode(gin.TestMode)

	mockUserRepo := new(mock.UserRepository)
	mockUserRepo.On("Update", mockUser).Return(nil).Once()

	router, rg := newMockRouter()
	api.NewUserAPI(rg, mockUserRepo)

	user_json, _ := json.Marshal(mockUser)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/users", bytes.NewReader(user_json))
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestDelete(t *testing.T) {
	mockUser := model.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	gin.SetMode(gin.TestMode)

	mockUserRepo := new(mock.UserRepository)
	mockUserRepo.On("Delete", mockUser.ID).Return(nil).Once()

	router, rg := newMockRouter()
	api.NewUserAPI(rg, mockUserRepo)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Code)
}
