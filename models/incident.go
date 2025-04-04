package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Constantes para los valores de status
const (
	StatusPendiente = "pendiente"
	StatusEnProceso = "en proceso"
	StatusResuelto  = "resuelto"
)

// Incident representa un reporte de incidente técnico
type Incident struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Reporter    string             `json:"reporter" bson:"reporter" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required,min=10"`
	Status      string             `json:"status" bson:"status"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// IncidentUpdate es una estructura para actualizar solo el campo status
type IncidentUpdate struct {
	Status string `json:"status" binding:"required"`
}

// ValidateStatus verifica si el status es válido
func ValidateStatus(status string) bool {
	return status == StatusPendiente || status == StatusEnProceso || status == StatusResuelto
}