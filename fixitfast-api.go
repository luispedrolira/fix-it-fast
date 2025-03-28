package main

import (
	"net/http"
	//"strconv"
	"github.com/gin-gonic/gin"

	//"errors"
)


type incident struct {
	ID    		int    	`json:"id"`
	Reporter 	string	`json:"reporter"`
	Description string 	`json:"description"`
	Status 		string 	`json:"status"`
	CreatedAt 	string 	`json:"created_at"`
}

var incidents = []incident{
	{ID: 1, Reporter: "John Doe", Description: "Se cayo el servidor", Status: "Pendiente", CreatedAt: "2025-01-01"},
	{ID: 2, Reporter: "Jane Doe", Description: "No puedo acceder a mi correo", Status: "En Curso", CreatedAt: "2025-01-02"},
	{ID: 3, Reporter: "John Doe", Description: "Mi computadora tiene malware :/", Status: "Resuelto", CreatedAt: "2025-01-03"},
}


// API Methods ↓
func getIncidents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, incidents)
}

func createIncident(c *gin.Context) {
	var newIncident incident

	if err := c.BindJSON(&newIncident); err != nil {
		return
	}

	incidents = append(incidents, newIncident)
	c.IndentedJSON(http.StatusCreated, newIncident)

}


func main(){
	router := gin.Default()
	router.GET("/incidents", getIncidents)
	router.POST("/incidents", createIncident)
	router.Run("localhost:8080")
}