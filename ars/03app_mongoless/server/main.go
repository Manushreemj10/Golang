package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Flight struct {
	Id           string
	Number       string
	AirlineName  string
	Source       string
	Destionation string
	Capacity     int
	Price        float32
}

func readAllFlights(c *gin.Context) {
	flights := []Flight{
		{Id: "1", Number: "AI 672", AirlineName: "Air India", Source: "Abu Dhabi", Destionation: "Mumbai", Capacity: 100, Price: 15000},
		{Id: "2", Number: "AI 673", AirlineName: "Air India", Source: "Abu Dhabi", Destionation: "Mumbai", Capacity: 100, Price: 15000},
	}
	c.JSON(http.StatusOK, flights)
}

func main() {
	//router
	r := gin.Default()

	//routes | API Endpoints
	r.GET("/flights", readAllFlights)

	//server (default port 8080) or manully r.Run(":9000")
	r.Run()

}
