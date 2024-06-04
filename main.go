package main

import (
	"GO/router"
	"GO/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	app         *gin.Engine
	mongoClient *mongo.Client
}

func NewServer() *Server {
	server := &Server{
		app: gin.Default(),
	}
	server.setNotFoundHandler()
	server.ErrorHandler()
	server.setConfiguration()
	server.setRouter()

	return server
}

func (s *Server) setConfiguration() {
	// s.dbConnect()
	utils.ConnectDB()
	s.app.Use(gin.Logger())
}

func (s *Server) dbConnect() {
	mongoURI := "mongodb+srv://vg100:vg100@cluster0.bszog.mongodb.net/test"
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	s.mongoClient = client
	fmt.Println("Connected to MongoDB!")
}

func (s *Server) setRouter() {
	router.AlbumRouter(s.app.Group("/api/album"))
	router.UserRouter(s.app.Group("/api/user"))
	s.app.GET("/health", HealthCheck)
}

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "API is up and working fine",
	})
}

func (s *Server) ErrorHandler() {
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

func (s *Server) setNotFoundHandler() {
	s.app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
		})
	})
}

func (s *Server) Run() {

	// err := godotenv.Load()

	port := ":8080"
	log.Printf("Server is connected to port %s", port)
	log.Fatal(s.app.Run(port))
}

type Student struct {
	Name  string
	Class string
}

func main() {
	// serverInstance := NewServer()
	// serverInstance.Run()

	students := []Student{
		{
			Name:  "Vijay",
			Class: "5th",
		},
	}

	student := []map[int]interface{}{
		{
			1: "Vijay",
			2: "5th",
		},
	}

	// arr := []interface{}{"name", 5, "uiu", 7, "jjj"}
	fmt.Println(students[0].Class)
	fmt.Println(student[0])

}
