package apperror

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

// AppError - Error details struct
type AppError struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Code    int    `json:"code"`    // application-specific error code
	Message string `json:"message"` // application-level error message
}

func create(code int, message string, httpStatus int) *AppError {
	return &AppError{
		HTTPStatusCode: httpStatus,
		Code:           code,
		Message:        message,
	}
}

func createWithErr(code int, message string, httpStatus int, err error) *AppError {
	return &AppError{
		Err:            err,
		HTTPStatusCode: httpStatus,
		Code:           code,
		Message:        message,
	}
}

// Error - Cast type to error
func (e *AppError) Error() string {
	return e.Message
}

// Render - Set response status code
func (e *AppError) Render(w http.ResponseWriter, r *http.Request) error {
	if e.Err != nil {
		log.Printf("Error: %v", e.Err.Error())
	}
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// General Errors
var (
	ErrNotFound     = create(10, "Resource not found", http.StatusNotFound)
	ErrUnauthorized = create(11, "Unauthorized", http.StatusUnauthorized)
)

// ErrRender - Render error
func ErrRender(err error) render.Renderer {
	if err, ok := err.(*AppError); ok {
		return err
	}
	return createWithErr(7, "Something went wrong", http.StatusInternalServerError, err)
}
