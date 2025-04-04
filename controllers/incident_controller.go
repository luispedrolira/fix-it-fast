package controllers

import (
	"incidentapi/models"
	"incidentapi/services"
	"incidentapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncidentController maneja las solicitudes HTTP para los incidentes
type IncidentController struct {
	service *services.IncidentService
}

// NewIncidentController crea un nuevo controlador de incidentes
func NewIncidentController(service *services.IncidentService) *IncidentController {
	return &IncidentController{
		service: service,
	}
}

// CreateIncident maneja la creación de un nuevo incidente
func (c *IncidentController) CreateIncident(ctx *gin.Context) {
	var incident models.Incident

	// Validar datos de entrada
	if err := ctx.ShouldBindJSON(&incident); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err))
		return
	}

	// Crear el incidente
	id, err := c.service.CreateIncident(ctx, &incident)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse("Error al crear incidente", err))
		return
	}

	incident.ID = id
	ctx.IndentedJSON(http.StatusCreated, utils.SuccessResponse("Incidente creado con éxito", incident))
}

// GetAllIncidents obtiene todos los incidentes
func (c *IncidentController) GetAllIncidents(ctx *gin.Context) {
	incidents, err := c.service.GetAllIncidents(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, utils.ErrorResponse("Error al obtener incidentes", err))
		return
	}

	ctx.IndentedJSON(http.StatusOK, utils.SuccessResponse("Incidentes obtenidos con éxito", incidents))
}

// GetIncidentByID obtiene un incidente por su ID
func (c *IncidentController) GetIncidentByID(ctx *gin.Context) {
	// Obtener y validar el ID
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Buscar el incidente
	incident, err := c.service.GetIncidentByID(ctx, id)
	if err == services.ErrNotFound {
		ctx.IndentedJSON(http.StatusNotFound, utils.ErrorResponse("Incidente no encontrado", err))
		return
	} else if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, utils.ErrorResponse("Error al obtener incidente", err))
		return
	}

	ctx.IndentedJSON(http.StatusOK, utils.SuccessResponse("Incidente obtenido con éxito", incident))
}

// UpdateIncidentStatus actualiza el estado de un incidente
func (c *IncidentController) UpdateIncidentStatus(ctx *gin.Context) {
	// Obtener y validar el ID
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Validar datos de entrada
	var update models.IncidentUpdate
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err))
		return
	}

	// Actualizar el estado
	err = c.service.UpdateIncidentStatus(ctx, id, update.Status)
	if err == services.ErrNotFound {
		ctx.IndentedJSON(http.StatusNotFound, utils.ErrorResponse("Incidente no encontrado", err))
		return
	} else if err == services.ErrInvalidStatus {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse("Estado inválido", err))
		return
	} else if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, utils.ErrorResponse("Error al actualizar incidente", err))
		return
	}

	ctx.IndentedJSON(http.StatusOK, utils.SuccessResponse("Estado de incidente actualizado con éxito", nil))
}

// DeleteIncident elimina un incidente
func (c *IncidentController) DeleteIncident(ctx *gin.Context) {
	// Obtener y validar el ID
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Eliminar el incidente
	err = c.service.DeleteIncident(ctx, id)
	if err == services.ErrNotFound {
		ctx.IndentedJSON(http.StatusNotFound, utils.ErrorResponse("Incidente no encontrado", err))
		return
	} else if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, utils.ErrorResponse("Error al eliminar incidente", err))
		return
	}

	ctx.IndentedJSON(http.StatusOK, utils.SuccessResponse("Incidente eliminado con éxito", nil))
}