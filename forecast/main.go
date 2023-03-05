package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const API_PREFIX = "https://pro.openweathermap.org/data/2.5/forecast/hourly?"

type Query struct {
	id    uint64
	appid string
}

func getForecast(c *gin.Context) {
	forecastData := make([]int, 0)
	c.IndentedJSON(http.StatusOK, forecastData)
}

func main() {
	router := gin.Default()
	router.GET("/forecast", getForecast)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
