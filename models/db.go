package models

import "github.com/leozz37/cartesian/services/db"

func CreateCoordinate(coordinate Coordinate) error {
	err := db.MySQL.Create(&coordinate).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCoordinates() error {
	err := db.MySQL.Delete(Coordinate{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FindCoordinates() ([]Coordinate, error) {
	var coordinates []Coordinate

	err := db.MySQL.Find(&coordinates).Error
	if err != nil {
		return nil, err
	}
	return coordinates, nil
}
