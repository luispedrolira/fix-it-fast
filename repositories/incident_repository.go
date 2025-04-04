package repositories

import (
	"context"
	"incidentapi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IncidentRepository maneja las operaciones de base de datos para los incidentes
type IncidentRepository struct {
	collection *mongo.Collection
}

// NewIncidentRepository crea un nuevo repositorio de incidentes
func NewIncidentRepository(client *mongo.Client, dbName, collName string) *IncidentRepository {
	collection := client.Database(dbName).Collection(collName)
	return &IncidentRepository{
		collection: collection,
	}
}

// Create inserta un nuevo incidente en la base de datos
func (r *IncidentRepository) Create(ctx context.Context, incident *models.Incident) (primitive.ObjectID, error) {
	// Establecer valores por defecto
	incident.CreatedAt = time.Now()
	if incident.Status == "" {
		incident.Status = models.StatusPendiente
	}

	// Insertar en MongoDB
	result, err := r.collection.InsertOne(ctx, incident)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Devolver el ID generado
	return result.InsertedID.(primitive.ObjectID), nil
}

// FindAll recupera todos los incidentes
func (r *IncidentRepository) FindAll(ctx context.Context) ([]models.Incident, error) {
	var incidents []models.Incident

	// Configurar la ordenación por fecha de creación (descendente)
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	// Ejecutar la consulta
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decodificar los resultados
	if err := cursor.All(ctx, &incidents); err != nil {
		return nil, err
	}

	return incidents, nil
}

// FindByID busca un incidente por su ID
func (r *IncidentRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Incident, error) {
	var incident models.Incident

	// Buscar por ID
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&incident)
	if err != nil {
		return nil, err
	}

	return &incident, nil
}

// UpdateStatus actualiza solo el campo status de un incidente
func (r *IncidentRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	// Definir la actualización para solo el campo status
	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}

	// Ejecutar la actualización
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

// Delete elimina un incidente por su ID
func (r *IncidentRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}