package ResponseHandler

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ResponseStatus bool        `json:"response_status"`
	HttpStatus     int         `json:"http_status"`
	Message        string      `json:"response_message"`
	Data           interface{} `json:"respose_data"`
}

var ginContext *gin.Context

func Go(c *gin.Context) Response {

	ginContext = c
	argResponse := Response{
		ResponseStatus: true,
		HttpStatus:     200,
		Message:        "Successfully Proced",
		Data:           nil,
	}

	return argResponse
}

func (response Response) SetHttpStatus(argHttpStatus int) Response {
	response.HttpStatus = argHttpStatus
	return response
}

func (response Response) SetMessage(argMessage string) Response {
	response.Message = argMessage
	return response
}

func (response Response) SetData(argData interface{}) Response {
	response.Data = argData
	return response
}

func (response Response) Apply() Response {
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) StopProcess() {
	panic(response.Message)
}

func (response Response) BadRequest(argMessage string) Response {
	response.Data = ""
	response.Message = argMessage
	response.HttpStatus = 400
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) DatabaseFetchFail(argDBError error) Response {
	dbError := make(map[string]interface{})
	dbError["showing_db_error"] = true
	if os.Getenv("ShowDatabaseErrorString") == "true" {

		dbError["database_error_msg"] = fmt.Sprintf("%v", argDBError)
		response.Data = dbError
	} else {
		dbError["showing_db_error"] = false
		response.Data = dbError
	}

	response.Message = "Database process failed"
	response.ResponseStatus = false
	response.HttpStatus = 500
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) DataNotFound(paramSearch string) Response {
	response.Message = "Data : " + paramSearch + " can't be found on database"
	response.ResponseStatus = false
	response.HttpStatus = 404
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) EnumNotFound(paramRequest string, enumAccepted []string) Response {
	response.Message = "Request : " + paramRequest + " not valid, accepted value : " + strings.Join(enumAccepted, ", ")
	response.ResponseStatus = false
	response.HttpStatus = 400
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) RequestUnmarshalFailure(unmarshalError string) Response {
	response.Message = unmarshalError
	response.ResponseStatus = false
	response.HttpStatus = 400
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) RequestValidationFailure(failResponse map[string]string) Response {

	response.Message = "Invalid payloads"
	response.ResponseStatus = false
	response.Data = map[string]interface{}{
		"validation_error": failResponse,
	}
	response.HttpStatus = 400
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) CustomProccessFailure(failResponse string) Response {

	response.Message = "Invalid payloads"
	response.ResponseStatus = false
	response.Data = map[string]interface{}{
		"process_failure": failResponse,
	}
	response.HttpStatus = 400
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) RequestJSONRequired(failResponse error) Response {

	response.Message = "Payload must be JSON format"
	response.ResponseStatus = false
	response.Data = map[string]interface{}{
		"validation_error": failResponse.Error(),
	}
	response.HttpStatus = 400
	ginContext.JSON(response.HttpStatus, response)
	return response
}
