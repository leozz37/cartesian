package handler

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leozz37/cartesian/models"
	"github.com/leozz37/cartesian/services/metrics"
)

var (
	reqAPIPointsCounter = *metrics.CreateCounter(
		"api_http_api_points_request_total",
		"Total http request on /api/points",
	)
)

// FindDistances GET /api/points
func FindDistances(c *gin.Context) {
	// Prometheus counter
	metrics.IncCounter(reqAPIPointsCounter)

	distance := c.Query("distance")
	x := c.Query("x")
	y := c.Query("y")

	err := validateQuery(x, y, distance)
	if err != nil {
		c.JSON(ResponseMessage(http.StatusBadRequest, err.Error()))
		return
	}

	// Converting string values to float
	distanceFloat, _ := strconv.ParseFloat(distance, 64)
	xFloat, _ := strconv.ParseFloat(x, 64)
	yFloat, _ := strconv.ParseFloat(y, 64)

	// Getting machinting distances
	coordinate := models.Coordinate{X: xFloat, Y: yFloat}
	matches, err := getWithinDistances(coordinate, distanceFloat)
	if err != nil {
		c.JSON(ResponseMessage(http.StatusBadRequest, err.Error()))
	}

	c.JSON(http.StatusOK, gin.H{"message": matches})
}

func getWithinDistances(coordinate models.Coordinate, distance float64) ([]models.Coordinate, error) {
	var matchingDistances []models.Coordinate

	// Getting coordinates from file
	coordinates, err := models.FindCoordinates()
	if err != nil {
		return nil, err
	}

	// Comparing the manhattan distance with the coordinates
	// from the file
	for _, point := range coordinates {
		manhattanDistance := models.CalculateManhattanDistance(coordinate, point)
		if manhattanDistance == distance {
			matchingDistances = append(matchingDistances, point)
		}
	}

	return matchingDistances, nil
}

func validateQueryValue(value string) error {
	match, _ := regexp.MatchString("[+-]?([0-9]*[.])?[0-9]+", value)
	if !match {
		return errors.New(value + " is an invalid value")
	}
	return nil
}

func validateQuery(x, y, distance string) error {
	err := validateQueryValue(distance)
	if err != nil {
		return err
	}
	err = validateQueryValue(x)
	if err != nil {
		return err
	}
	err = validateQueryValue(y)
	if err != nil {
		return err
	}
	return nil
}
