package services

import (
	"context"
	"errors"
	"incidentapi/models"
	"incidentapi/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Errores del servicio
var (
	ErrInvalidStatus = errors.New("estado no válido, debe ser 'pendiente', 'en proceso' o 'resuelto'")
	ErrNotFound      = errors.New("incidente no encontrado")
)

// IncidentService implementa la lógica de negocio para los incidentes
type IncidentService struct {
	repo *repositories.IncidentRepository
}

// NewIncidentService crea un nuevo servicio de incidentes
func NewIncidentService(repo *repositories.IncidentRepository) *IncidentService {
	return &IncidentService{
		repo: repo,
	}
}

// CreateIncident crea un nuevo incidente
func (s *IncidentService) CreateIncident(ctx context.Context, incident *models.Incident) (primitive.ObjectID, error) {
	// Validar el estado si se proporciona
	if incident.Status != "" && !models.ValidateStatus(incident.Status) {
		return primitive.NilObjectID, ErrInvalidStatus
	}

	// Crear el incidente
	return s.repo.Create(ctx, incident)
}

// GetAllIncidents obtiene todos los incidentes
func (s *IncidentService) GetAllIncidents(ctx context.Context) ([]models.Incident, error) {
	return s.repo.FindAll(ctx)
}

// GetIncidentByID obtiene un incidente por su ID
func (s *IncidentService) GetIncidentByID(ctx context.Context, id primitive.ObjectID) (*models.Incident, error) {
	incident, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return incident, nil
}

// UpdateIncidentStatus actualiza solo el campo status de un incidente
func (s *IncidentService) UpdateIncidentStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	// Validar que el incidente existe
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	// Validar el estado
	if !models.ValidateStatus(status) {
		return ErrInvalidStatus
	}

	// Actualizar el estado
	return s.repo.UpdateStatus(ctx, id, status)
}

// DeleteIncident elimina un incidente
func (s *IncidentService) DeleteIncident(ctx context.Context, id primitive.ObjectID) error {
	// Validar que el incidente existe
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	// Eliminar el incidente
	return s.repo.Delete(ctx, id)
}