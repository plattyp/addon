package endpoints

import (
	"fmt"
	"net/http"

	validator "gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
	"github.com/plattyp/addon/db"
)

// Endpointer holds all context for the endpoint to use
type Endpointer struct {
	databaser *db.Databaser
}

// FieldError allows us to return structured errors to the webhook response
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewEndpointer returns a new endpointer to be used
func NewEndpointer(d *db.Databaser) *Endpointer {
	return &Endpointer{databaser: d}
}

// PrintMessage returns back the structured message
func (f FieldError) PrintMessage() string {
	return "[" + f.Field + "] " + f.Message
}

// Success returns generic success message
func Success(message string, c *gin.Context) {
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  true,
			"message": message,
		},
	)
}

// Error returns generic error message
func Error(e string, c *gin.Context) {
	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"status":  false,
			"message": e,
		},
	)
}

// ValidationError returns a validation error message
func ValidationError(f FieldError, c *gin.Context) {
	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"status":  false,
			"message": f.PrintMessage(),
		},
	)
}

// HandleErrors formats binding errors to a structured slice of FieldErrors
func HandleError(err error) FieldError {
	var error FieldError
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		for _, valErr := range errs {
			error = FieldError{
				Field:   valErr.FieldNamespace,
				Message: fmt.Sprintf("Validation error due to the following tag '%s'", valErr.Tag),
			}
		}
	} else {
		error = FieldError{
			Field:   "Generic",
			Message: err.Error(),
		}
	}

	return error
}
