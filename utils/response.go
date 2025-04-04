package utils

// Response es la estructura base para todas las respuestas HTTP
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// SuccessResponse crea una respuesta exitosa
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse crea una respuesta de error
func ErrorResponse(message string, err error) Response {
	var errorMsg interface{}
	if err != nil {
		errorMsg = err.Error()
	}

	return Response{
		Success: false,
		Message: message,
		Error:   errorMsg,
	}
}