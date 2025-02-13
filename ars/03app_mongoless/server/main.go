package main

import (
	//"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Flight struct {
	Id           string  `json:"id"`
	Number       string	 `json:"number"`
	AirlineName  string  `json:"airline_name"`
	Source       string  `json:"source"`
	Destionation string  `json:"destination"`
	Capacity     int     `json:"capacity"`
	Price        float32 `json:"price"`
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

	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//routes | API Endpoints
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightById)
	r.POST("/flights", createFlight)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)


	//server (default port 8080) or manully r.Run(":9000")
	r.Run(":8080")

}
