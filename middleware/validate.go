package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// Shareable Errors
var (
	ErrorInternalError = errors.New("Something went wrong")
)

// ValidationErrorToText translates a field error to more human readable
func ValidationErrorToText(inputField string, e *validator.FieldError) string {
	switch e.Tag {
	case "required":
		return fmt.Sprintf("%s is required", inputField)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", inputField, e.Param)
	case "min":
		return fmt.Sprintf("%s must be longer than %s", inputField, e.Param)
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", inputField, e.Param)
	}
	return fmt.Sprintf("%s is not valid", inputField)
}

// ValidateErrors collects all errors from binding the Request to structs and returns if any are present
func ValidateErrors() gin.HandlerFunc {
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
					resource, _ := c.Get("resource")
					errs, ok := e.Err.(validator.ValidationErrors)
					if !ok {
						c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Please reevaluate your request parameters"})
						return
					}
					errorMessages := []string{}
					for _, err := range errs {
						inputField := err.Field
						if resource != nil {
							field, _ := reflect.TypeOf(resource).Elem().FieldByName(err.Field)
							inputField = field.Tag.Get("json")
						}
						errorMessages = append(errorMessages, ValidationErrorToText(inputField, err))
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
