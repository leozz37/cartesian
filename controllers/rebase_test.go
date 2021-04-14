package controllers

import (
	"os"
	"testing"

	"github.com/leozz37/cartesian/models"
	"github.com/leozz37/cartesian/services/db"
)

func setupMySQL() {
	db.ConnectMySQL("sqlite3", "testing.db")
	models.DeleteCoordinates()
	db.AutoMigration(models.Coordinate{})
}

func TestRebaseSuccess(t *testing.T) {
	setupMySQL()
	os.Setenv("DATA_PATH", "../data/points.json")

	Rebase()
	result, err := models.FindCoordinates()
	if err != nil {
		t.Error(err)
	}

	expectedCoordinate := models.Coordinate{X: 63, Y: -72}

	if result[0].X != expectedCoordinate.X {
		t.Error("X coordinate is different")
	} else if result[0].Y != expectedCoordinate.Y {
		t.Error("Y coordinate is different")
	}
}

func TestRebaseFailFilePath(t *testing.T) {
	setupMySQL()
	os.Setenv("DATA_PATH", "random/path/unknow.json")

	// Checking if Rebase() will panic!
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Rebase()
}
