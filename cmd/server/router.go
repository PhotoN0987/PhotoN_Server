package server

import (
	"net/http"
	"photon-server/pkg/api"
	"photon-server/pkg/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	// すべてのアクセス許可 & Content-Disposition許可
	config := cors.Config{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Disposition"},
	}
	router.Use(cors.New(config))
	router.StaticFS("/web", http.Dir("web"))

	return router
}

// SetUpRouter Setup all api routing
func (s *Server) SetUpRouter() *gin.Engine {
	// Group v1
	apiV1 := s.router.Group("api/v1")
	s.userRoutes(apiV1)
	s.fileRoutes(apiV1)
	return s.router
}

func (s *Server) userRoutes(rg *gin.RouterGroup) {
	repository := repository.NewUserRepository(s.db)
	api.NewUserAPI(rg, repository)
}

func (s *Server) fileRoutes(rg *gin.RouterGroup) {
	api.NewFileAPI(rg)
}
