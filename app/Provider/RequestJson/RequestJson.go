package RequestJson

import (
	"fmt"
	"inventori/app/Provider/ResponseHandler"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var MustJSON bool

func Validate(argPayloadJSON error, c *gin.Context) bool {
	failResponse := make(map[string]string)
	
	fmt.Println("argPayloadJSON", MustJSON)
	if argPayloadJSON != nil {
		if MustJSON {
			if reflect.TypeOf(argPayloadJSON).String() == "*json.SyntaxError" {
				ResponseHandler.Go(c).RequestJSONRequired(argPayloadJSON)
				return true
			}
		}
		if reflect.TypeOf(argPayloadJSON).String() == "*json.UnmarshalTypeError" {
			ResponseHandler.Go(c).RequestUnmarshalFailure(argPayloadJSON.Error())
		} else {
			for _, e := range argPayloadJSON.(validator.ValidationErrors) {
				field := strings.ToLower(e.Field())
				errorMessage := fmt.Sprintf("%s field is %s", e.Field(), e.ActualTag())
				failResponse[field] = errorMessage
			}
			ResponseHandler.Go(c).RequestValidationFailure(failResponse)
		}

		return true
	}
	return false
}
