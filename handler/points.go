package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leozz37/cartesian/models"
	"github.com/leozz37/cartesian/services/metrics"
)

var (
	reqPointsCounter = *metrics.CreateCounter(
		"api_http_points_request_total",
		"Total http request on /points",
	)
)

// GET /points
func GetPoints(c *gin.Context) {
	// Prometheus counter
	metrics.IncCounter(reqPointsCounter)

	coordinates, err := models.FindCoordinates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": coordinates})
}
