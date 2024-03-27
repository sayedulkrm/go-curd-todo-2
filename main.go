package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sayedulkrm/go-curd-todo-2/database"
	"github.com/sayedulkrm/go-curd-todo-2/lib"
	"github.com/sayedulkrm/go-curd-todo-2/routes"
	"github.com/sirupsen/logrus"
)

func main() {

	// Configure logging
	lib.ConfigureLogger()

	// Set up routes and middleware
	root := routes.SetupRoutes()

	// orgin := os.Getenv("FRONTEND_URL")

	// fmt.Println(orgin)

	// Create CORS handler with custom options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://example.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		// Debug:            true,
		// Add other custom options as needed
	}

	c := cors.New(corsOptions)

	// ======================================

	// Wrap your existing handler with CORS middleware
	handler := c.Handler(root)

	// Start the server
	startServer(handler)

}

func startServer(handler http.Handler) {
	// Start the server

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}

	// Start DB
	database.DBInstance()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8000"
	}

	logrus.Warn("Server running on", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		logrus.Fatal("Failed to start server: ", err)
	}
}
