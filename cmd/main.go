package main

import (
	"context"
	"incidentapi/config"
	"incidentapi/routes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Configurar cliente MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conexión a MongoDB
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	// Verificar la conexión
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("No se pudo conectar a MongoDB:", err)
	}
	log.Println("Conectado exitosamente a MongoDB")

	// Inicializar el router de Gin
	router := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(router, client)

	// Iniciar el servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}