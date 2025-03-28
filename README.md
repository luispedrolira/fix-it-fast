# FixItFast API
Este proyecto es una API simple desarrollada en Go utilizando el framework Gin para manejar reportes de incidentes.

## Rutas de la API

La API proporciona las siguientes rutas:

### Obtener todos los incidentes
- **Método:** `GET`
- **Ruta:** `/incidents`
- **Descripción:** Devuelve una lista de todos los incidentes reportados.

### Crear un nuevo incidente
- **Método:** `POST`
- **Ruta:** `/incidents`
- **Descripción:** Permite crear un nuevo incidente enviando un JSON con los datos.
- **Ejemplo de request:**
  ```json
  {
    "id": 4,
    "reporter": "Alice Doe",
    "description": "Problema con la red",
    "status": "Pendiente",
    "created_at": "2025-01-04"
  }
  ```

## Cómo correr la API

1. **Clona el repositorio**
   ```sh
   git clone <URL_DEL_REPOSITORIO>
   cd <NOMBRE_DEL_PROYECTO>
   ```

2. **Ejecuta la API**
   ```sh
   go run fixitfast-api.go
   ```

4. **Prueba la API**
   - Abre un navegador o usa `curl` para probar:
     ```sh
     curl http://localhost:8080/incidents
     ```
   - También puedes usar herramientas como Postman o HTTPie.

## Tecnologías utilizadas
- Go
- Gin (framework para manejo de rutas y middleware)

---
¡Gracias por usar la FixItFast API!

