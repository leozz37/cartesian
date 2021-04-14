package routes

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leozz37/cartesian/handler"
)

func InitRoutes() {
	log.Println("Starting routes")
	r := gin.Default()

	// Setting log format
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.GET("/api/points", handler.FindDistances)
	r.GET("points", handler.GetPoints)
	r.NoRoute(noRoute)

	r.Run()
}

func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Page not found"})
}
