package ErrorHandler

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/withmandala/go-log"
)

type SetError struct {
	ErrorValue error
	checkName  string
}

func Err(argErrorValue error) SetError {
	var setERror SetError
	setERror.ErrorValue = argErrorValue
	return setERror
}

func (value SetError) Check(argCheckName string) SetError {
	value.checkName = argCheckName
	return value
}

func (value SetError) Fatal() {
	logger := log.New(os.Stderr)
	if value.ErrorValue != nil {
		logger.Fatal(value.ErrorValue)
	}
	logger.Info("Success Processing", value.checkName)
}

func (value SetError) Error() {
	logger := log.New(os.Stderr)
	logger.Info("Error When : ", value.checkName)
	if value.ErrorValue != nil {
		logger.Error(value.ErrorValue)
	}
	panic("Validation Error")
}

func (value SetError) ErrorResponse(c *gin.Context) {
	if value.ErrorValue != nil {
		fmt.Println("value.ErrorValue", value.ErrorValue)
		c.JSON(500, gin.H{
			"response_status":  false,
			"response_message": "Payload is not json",
			"respose_data":     value.ErrorValue,
		})
		// panic(value.checkName)
	}
}
