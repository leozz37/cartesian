package controllers

import (
	"log"
	"os"

	"github.com/leozz37/cartesian/models"
)

// Rebase delete as Coordinates from Database, read from the
// JSON file save them on Database
func Rebase() {
	log.Println("Migrating points from JSON file to database")
	models.DeleteCoordinates()

	filePath := os.Getenv("DATA_PATH")

	coordinates, err := models.GetCoordinatesFromFile(filePath)
	if err != nil {
		panic(err)
	}

	models.CreateCoordinates(coordinates)
}
