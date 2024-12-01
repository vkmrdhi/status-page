package main

import (
	"backend/config"
	"backend/routes"
	"log"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Setup database connection
	err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Setup routes
	r := routes.SetupRouter()

	// Start the server
	r.Run(":8080")
}
