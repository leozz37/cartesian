package models

import (
	"testing"
)

func TestCreateCoordinate(t *testing.T) {
	setupMySQL()

	coordinate := Coordinate{X: 63, Y: -72}
	CreateCoordinate(coordinate)

	result, err := FindCoordinates()
	if err != nil {
		t.Error(err)
	}

	if result[0].X != coordinate.X {
		t.Error("X coordinate is different")
	} else if result[0].Y != coordinate.Y {
		t.Error("Y coordinate is different")
	}
}

func TestDeleteCoordinates(t *testing.T) {
	setupMySQL()

	coordinates, err := GetCoordinatesFromFile("../data/points.json")
	if err != nil {
		t.Error(err)
	}

	CreateCoordinates(coordinates)
	DeleteCoordinates()

	result, err := FindCoordinates()
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Error("Coordinates was not deletes")
	}
}

func TestFindCoordinates(t *testing.T) {
	setupMySQL()

	coordinates, err := GetCoordinatesFromFile("../data/points.json")
	if err != nil {
		t.Error(err)
	}
	CreateCoordinates(coordinates)

	result, err := FindCoordinates()
	if err != nil {
		t.Error(err)
	}

	if len(result) != 100 {
		t.Error("Coordinates was not created")
	}
}
