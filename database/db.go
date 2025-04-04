package database

import (
	"context"
	"time"

	"example/fix-it-fast/config"
	"example/fix-it-fast/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Colección en MongoDB
var collection = config.DB.Collection("incidents")

// Obtener todos los incidentes
func GetIncidents() ([]models.Incident, error) {
	var incidents []models.Incident
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var incident models.Incident
		cursor.Decode(&incident)
		incidents = append(incidents, incident)
	}
	return incidents, nil
}

// Obtener un incidente por ID
func GetIncidentByID(id string) (models.Incident, error) {
	var incident models.Incident
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return incident, err
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&incident)
	return incident, err
}

// Crear un nuevo incidente
func CreateIncident(incident models.Incident) error {
	incident.ID = primitive.NewObjectID()
	incident.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), incident)
	return err
}

// Actualizar el estado de un incidente
func UpdateIncidentStatus(id string, status string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"status": status}})
	return err
}

// Eliminar un incidente
func DeleteIncident(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
