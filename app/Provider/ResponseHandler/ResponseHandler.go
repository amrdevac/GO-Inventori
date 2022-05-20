package ResponseHandler

import "github.com/gin-gonic/gin"

type Response struct {
	HttpStatus int         `json:"httpStatus"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

var ginContext *gin.Context

func Go(c *gin.Context) Response {

	ginContext = c
	argResponse := Response{
		HttpStatus: 200,
		Message:    "Successfully Proced",
		Data:       nil,
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

func (response Response) Get() Response {
	ginContext.JSON(response.HttpStatus, response)
	return response
}

func (response Response) StopProcess() {
	panic(response.Message)
}
