package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pizza struct {
	Id       string
	Name     string
	Size     int
	Price    float32
	Category string
}

func readAllPizza(c *gin.Context) {
	pizzaes := []Pizza{
		{Id: "1", Name: "Corn Pizza", Category: "Indian", Price: 200},
		{Id: "2", Name: "Papad Pizza", Category: "Italian", Price: 150},
	}
	c.JSON(http.StatusOK, pizzaes)
}

func readPizzaById(c *gin.Context) {
	id := c.Param("id")
	pizzaes := Pizza{ Id: id, Name: "Corn Pizza", Category: "Indian", Price: 200}

	c.JSON(http.StatusOK, pizzaes)
}

func createPizza(c *gin.Context) {
	var jbodyFlight Pizza
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Error." + err.Error()})
		return
	}
	createdPizza := Pizza{Id: "4", Name: "Papad Pizza", Category: "Italian", Price: 150}

	c.JSON(http.StatusCreated,
		gin.H{"message": "Pizza Created Successfully.",
			"pizza": createdPizza})
}

func updatePizza(c *gin.Context) {
	id := c.Param("id")
	var jbodyFlight Pizza
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Error." + err.Error()})
		return
	}
	updatedPizza := Pizza{Id: id, Name: "Papad Pizza", Category: "Italian", Price: 150}

	c.JSON(http.StatusOK,
		gin.H{"message": "Pizza Updated Successfully.",
			"pizza": updatedPizza})
}

func deletePizza(c *gin.Context) {
	//id :=c.Param("id")

	c.JSON(http.StatusOK,
		gin.H{"message": "Pizza Updated Successfully."})
}

func main() {
	//router
	r := gin.Default()

	//routes | API Endpoints
	r.GET("/pizza", readAllPizza)
	r.GET("/pizza/:id", readPizzaById)
	r.POST("/pizza", createPizza)
	r.PUT("/pizza/:id", updatePizza)
	r.DELETE("/pizza/:id", deletePizza)

	//server (default port 8080) or manully r.Run(":9000")
	r.Run(":8080")

}
