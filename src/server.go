package src

import (
	"log"
	"net/http"

	"GO/src/router"
	"GO/src/utils"

	"github.com/gin-gonic/gin"
)

type ServerType struct {
	app *gin.Engine
}

func Server() *ServerType {
	server := &ServerType{
		app: gin.Default(),
	}
	server.setNotFoundHandler()
	server.ErrorHandler()
	server.setConfiguration()
	server.setRouter()

	return server
}

func (s *ServerType) setConfiguration() {
	utils.ConnectDB()
	s.app.Use(gin.Logger())
}

func (s *ServerType) setRouter() {
	router.AlbumRouter(s.app.Group("/api/album"))
	router.UserRouter(s.app.Group("/api/user"))
	router.ProductRouter(s.app.Group("/api/product"))
	s.app.GET("/health", HealthCheck)
}

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "API is up and working fine",
	})
}

func (s *ServerType) ErrorHandler() {
	s.app.Use(func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	})
}

func (s *ServerType) setNotFoundHandler() {
	s.app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
		})
	})
}

func (s *ServerType) Run() {

	port := ":8080"
	log.Printf("Server is connected to port %s", port)
	log.Fatal(s.app.Run(port))
}
