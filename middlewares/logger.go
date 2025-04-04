package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware para registrar información sobre las solicitudes HTTP
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Tiempo de inicio
		startTime := time.Now()

		// Procesar la solicitud
		ctx.Next()

		// Tiempo de finalización
		endTime := time.Now()

		// Registrar la información
		latency := endTime.Sub(startTime)
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()

		// Usar el paquete log estándar en lugar de ctx.Logger()
		log.Printf("[%s] %s %s %d %v", 
			method, 
			path, 
			ctx.ClientIP(), 
			statusCode, 
			latency,
		)
	}
}