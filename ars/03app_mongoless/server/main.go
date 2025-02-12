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

func readFlightById(c *gin.Context) {
	id := c.Param("id")
	flights := Flight {Id: id, Number: "AI 672", AirlineName: "Air India", Source: "Abu Dhabi", Destionation: "Mumbai", Capacity: 100, Price: 15000}
	
	c.JSON(http.StatusOK, flights)
}

func createFlight(c *gin.Context){
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError , 
		gin.H{"error" : "Server Error." + err.Error()})
		return
	}
	createdFlight := Flight {Id: "11", Number: "AI624", AirlineName: "Air India", Source: "Abu Dhabi", Destionation: "Mumbai", Capacity: 100, Price: 15000}
	
	c.JSON(http.StatusCreated,
		gin.H{"message" : "Flight Created Successfully.",
			"flight" : createdFlight})
}

func updateFlight(c *gin.Context){
	id :=c.Param("id")
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError , 
		gin.H{"error" : "Server Error." + err.Error()})
		return
	}
	updatedFlight := Flight {Id: id, Number: "AI624", AirlineName: "Air India", Source: "Abu Dhabi", Destionation: "Mumbai", Capacity: 100, Price: 15000}
	
	c.JSON(http.StatusOK,
		gin.H{"message" : "Flight Updated Successfully.",
			"flight" : updatedFlight})
}

func deleteFlight(c *gin.Context){
	//id :=c.Param("id")
	
	c.JSON(http.StatusOK,
		gin.H{"message" : "Flight Updated Successfully."})
}

func main() {
	//router
	r := gin.Default()

	//routes | API Endpoints
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightById)
	r.POST("/flights", createFlight)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)


	//server (default port 8080) or manully r.Run(":9000")
	r.Run(":8080")

}
