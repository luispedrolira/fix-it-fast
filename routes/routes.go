package routes

import (
	"incidentapi/config"
	"incidentapi/controllers"
	"incidentapi/middlewares"
	"incidentapi/repositories"
	"incidentapi/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(router *gin.Engine, client *mongo.Client) {
	// Middlewares globales
	router.Use(middlewares.Logger())
	
	// Crear las dependencias
	repo := repositories.NewIncidentRepository(client, config.DatabaseName, config.CollectionName)
	service := services.NewIncidentService(repo)
	controller := controllers.NewIncidentController(service)

	// Grupo de rutas para API v1
	v1 := router.Group("/incidents")
	{
		// Endpoints para incidentes
		v1.POST("", controller.CreateIncident)
		v1.GET("", controller.GetAllIncidents)
		v1.GET("/:id", controller.GetIncidentByID)
		v1.PUT("/:id", controller.UpdateIncidentStatus)
		v1.DELETE("/:id", controller.DeleteIncident)
	}
}