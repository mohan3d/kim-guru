package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/mohan3d/apixu-go"
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
	apixuAPIKey := os.Getenv("APIXU_KEY")
	apixuClient := apixu.NewClient(apixuAPIKey)
	apixuClient.Current("Cairo")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/cairo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "kim.tmp.html", nil)
	})

	router.Run(":" + port)
}
