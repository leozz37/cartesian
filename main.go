package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/leozz37/cartesian/controllers"
	"github.com/leozz37/cartesian/models"
	"github.com/leozz37/cartesian/routes"
	"github.com/leozz37/cartesian/services/db"
	"github.com/leozz37/cartesian/services/metrics"
)

func main() {
	log.Println("STATING THE APPLICATION")
	log.Println("Reading the .env file")
	// Load environment variables from .env file
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init Database and Migrate coordinates model
	db.ConnectMySQL(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_DSN"))
	db.AutoMigration(models.Coordinate{})

	// Get data from JSON file and save to DB
	controllers.Rebase()

	// Init Prometheus
	go metrics.InitPrometheus()

	// Setting up routes
	routes.InitRoutes()
}
