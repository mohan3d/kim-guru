package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	apixu "github.com/mohan3d/apixu-go"
)

var apiKey = os.Getenv("APIXU_KEY")

type weatherInfo struct {
	TempC   float64
	TempF   float64
	Status  string
	Country string
	Region  string
}

func getWeatherInfo(city string) (*weatherInfo, error) {
	client := apixu.NewClient(apiKey)
	currentWeather, err := client.Current(city)

	if err != nil {
		return nil, err
	}

	return &weatherInfo{
		TempC:   currentWeather.Current.TempC,
		TempF:   currentWeather.Current.TempF,
		Status:  currentWeather.Current.Condition.Text,
		Country: currentWeather.Location.Country,
		Region:  currentWeather.Location.Region,
	}, nil
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "kim.tmpl.html", nil)
	})

	router.GET("/weather", func(c *gin.Context) {
		query := c.Request.URL.Query()
		lat := query.Get("lat")
		long := query.Get("long")
		q := fmt.Sprintf("%v,%v", lat, long)

		weather, err := getWeatherInfo(q)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, weather)
	})

	router.Run(":" + port)
}
