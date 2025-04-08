package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/adapters/inbound/http"
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/adapters/outbound/postgres"
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/application"
)

func main() {
	// Initialize the database connection
	db, err := postgres.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize the repositories
	userRepository := postgres.NewUserRepository(db)

	// Initialize the services
	userService := application.NewUserService(userRepository)

	// Initialize the handlers
	userHandler := http.NewUserHandler(userService)

	// Initialize the server
	r := gin.Default()

	// Set up routes
	r.Group("/users").
		GET("/", userHandler.GetAll).
		GET("/:id", userHandler.GetByID).
		POST("/", userHandler.Create).
		PUT("/:id", userHandler.UpdateByID).
		DELETE("/:id", userHandler.DeleteByID)

	// Start the server
	log.Fatalln(r.Run())
}
