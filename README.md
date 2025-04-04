# FixItFast API

Este proyecto es una API REST completa desarrollada en Go utilizando el framework Gin y MongoDB para gestionar reportes de incidentes técnicos dentro de una empresa.

## Estructura del Proyecto

```
.
├── cmd
│   └── main.go                  # Punto de entrada de la aplicación
├── config
│   └── config.go                # Configuración de la aplicación
├── controllers
│   └── incident_controller.go   # Controladores para los endpoints
├── middlewares
│   └── logger.go                # Middleware para logging
├── models
│   └── incident.go              # Modelo de datos para incidentes
├── repositories
│   └── incident_repository.go   # Capa de acceso a datos
├── routes
│   └── routes.go                # Definición de rutas
├── services
│   └── incident_service.go      # Lógica de negocio
├── utils
│   └── response.go              # Utilidades para respuestas HTTP
```

## Rutas de la API

La API proporciona las siguientes rutas:

### Obtener todos los incidentes
- **Método:** `GET`
- **Ruta:** `/incidents`
- **Descripción:** Devuelve una lista de todos los incidentes reportados.
- **Respuesta:**
  ```json
  {
    "success": true,
    "message": "Incidentes obtenidos con éxito",
    "data": [
      {
        "id": "6457a86c72e1c91234567890",
        "reporter": "Juan Pérez",
        "description": "La impresora no funciona correctamente desde ayer",
        "status": "pendiente",
        "created_at": "2025-04-03T14:30:45.123Z"
      },
      // más incidentes...
    ]
  }
  ```

### Obtener un incidente específico
- **Método:** `GET`
- **Ruta:** `/incidents/:id`
- **Descripción:** Devuelve un incidente específico por su ID.
- **Respuesta:**
  ```json
  {
    "success": true,
    "message": "Incidente obtenido con éxito",
    "data": {
      "id": "6457a86c72e1c91234567890",
      "reporter": "Juan Pérez",
      "description": "La impresora no funciona correctamente desde ayer",
      "status": "pendiente",
      "created_at": "2025-04-03T14:30:45.123Z"
    }
  }
  ```

### Crear un nuevo incidente
- **Método:** `POST`
- **Ruta:** `/incidents`
- **Descripción:** Permite crear un nuevo incidente enviando un JSON con los datos.
- **Ejemplo de request:**
  ```json
  {
    "reporter": "Alice Doe",
    "description": "Problema con la red del tercer piso",
    "status": "pendiente"
  }
  ```
- **Respuesta:**
  ```json
  {
    "success": true,
    "message": "Incidente creado con éxito",
    "data": {
      "id": "6457a86c72e1c91234567890",
      "reporter": "Alice Doe",
      "description": "Problema con la red del tercer piso",
      "status": "pendiente",
      "created_at": "2025-04-03T14:32:15.678Z"
    }
  }
  ```

### Actualizar el estado de un incidente
- **Método:** `PUT`
- **Ruta:** `/incidents/:id`
- **Descripción:** Actualiza únicamente el campo `status` de un incidente.
- **Ejemplo de request:**
  ```json
  {
    "status": "en proceso"
  }
  ```
- **Respuesta:**
  ```json
  {
    "success": true,
    "message": "Estado de incidente actualizado con éxito",
    "data": null
  }
  ```

### Eliminar un incidente
- **Método:** `DELETE`
- **Ruta:** `/incidents/:id`
- **Descripción:** Elimina un incidente por su ID.
- **Respuesta:**
  ```json
  {
    "success": true,
    "message": "Incidente eliminado con éxito",
    "data": null
  }
  ```

## Requisitos

- Go 1.16 o superior
- MongoDB 4.4 o superior
- Conexión a Internet para descargar dependencias

## Configuración

La configuración de la base de datos se encuentra en el archivo `config/config.go`. Por defecto, la aplicación se conectará a:

```go
const (
    MongoURI = "mongodb://localhost:27017"
    DatabaseName = "support"
    CollectionName = "incidents"
)
```

## Cómo instalar y ejecutar la API

1. **Clona el repositorio**
   ```sh
   git clone <URL_DEL_REPOSITORIO>
   cd <NOMBRE_DEL_PROYECTO>
   ```

2. **Instala las dependencias**
   ```sh
   go mod download
   ```

3. **Asegúrate de que MongoDB esté en ejecución**
   ```sh
   # Si usas Docker para MongoDB:
   docker run -d -p 27017:27017 --name mongodb mongo
   ```

4. **Ejecuta la API**
   ```sh
   go run cmd/main.go
   ```

5. **Prueba la API**
   - Abre un navegador o usa `curl` para probar:
     ```sh
     curl http://localhost:8080/incidents
     ```
   - También puedes usar herramientas como Postman o Thunder Client (VSCode).


---

¡Gracias por usar la FixItFast API!
