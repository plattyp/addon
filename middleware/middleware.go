package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// Shareable Errors
var (
	ErrorInternalError = errors.New("Something went wrong")
)

// ValidationErrorToText translates a field error to more human readable
func ValidationErrorToText(e *validator.FieldError) string {
	switch e.Tag {
	case "required":
		return fmt.Sprintf("%s is required", e.Field)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field, e.Param)
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field, e.Param)
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field, e.Param)
	}
	return fmt.Sprintf("%s is not valid", e.Field)
}

// Errors collects all errors and returns them
func Errors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Only run if there are some errors to handle
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Find out what type of error it is
				switch e.Type {
				case gin.ErrorTypePublic:
					// Only output public errors if nothing has been written yet
					if !c.Writer.Written() {
						c.JSON(c.Writer.Status(), gin.H{"success": false, "message": e.Error()})
					}
				case gin.ErrorTypeBind:
					errs := e.Err.(validator.ValidationErrors)
					errorMessages := []string{}
					for _, err := range errs {
						errorMessages = append(errorMessages, ValidationErrorToText(err))
					}

					// Make sure we maintain the preset response status
					status := http.StatusBadRequest
					if c.Writer.Status() != http.StatusOK {
						status = c.Writer.Status()
					}

					message := strings.Join(errorMessages, ", ")
					c.JSON(status, gin.H{"success": false, "message": message})
				}

			}
			// If there was no public or bind error, display default 500 message
			if !c.Writer.Written() {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": ErrorInternalError.Error()})
			}
		}
	}
}
