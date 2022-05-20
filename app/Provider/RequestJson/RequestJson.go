package RequestJson

import (
	"fmt"
	"login-sistem-jwt/app/Provider/ErrorHandler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(argPayloadJSON error, c *gin.Context) {
	ErrorHandler.Err(argPayloadJSON).Check("Payload not json").Error()
	failResponse := make(map[string]string)
	

	if argPayloadJSON != nil {
		for _, e := range argPayloadJSON.(validator.ValidationErrors) {
			field := strings.ToLower(e.Field())
			errorMessage := fmt.Sprintf("%s field is %s", e.Field(), e.ActualTag())

			failResponse[field] = errorMessage
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  failResponse,
			"status": http.StatusBadRequest,
		})
		panic("Validation Error")
	}

}
