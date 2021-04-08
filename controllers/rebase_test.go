package controllers

import (
	"testing"

	"github.com/leozz37/cartesian/models"
)

func TestGetCoordinatesFromFile(t *testing.T) {
	// The very first coordinate from points.json
	expectedPoint := models.Coordinate{X: 63, Y: -72}

	coordinates, err := getCoordinatesFromFile("../data/points.json")
	if err != nil {
		t.Error(err)
	}

	for _, v := range coordinates {
		if v.X == expectedPoint.X && v.Y == expectedPoint.Y {
			return
		}
	}

	t.Error("Couldn't find coordinates")
}
