package item

import (
	"fmt"
	"inventori/app/Provider/RequestJson"
	"inventori/app/Provider/ResponseHandler"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	RequestJson.MustJSON = true
	request := ListBarangRequest{}
	if RequestJson.Validate(c.ShouldBindJSON(&request), c) {
		return
	}

	listData, isSucccess := GetAll(request)
	if !isSucccess {
		return
	}

	countAll, isSucccess := CountAll(request)
	if !isSucccess {
		return
	}

	result := make(map[string]interface{})
	result["count"] = countAll
	result["data"] = listData

	ResponseHandler.Go(c).SetData(result).Apply()
}

func Detail(c *gin.Context) {
	RequestJson.MustJSON = true
	request := DetailItemRequest{}
	if RequestJson.Validate(c.ShouldBindJSON(&request), c) {
		return
	}

	listData, isSucccess := GetOnce(request)
	if !isSucccess {
		return
	}

	result := make(map[string]interface{})
	result["data"] = listData

	ResponseHandler.Go(c).SetData(result).Apply()
}

func Store(c *gin.Context) {

	request := ItemTable{}
	if RequestJson.Validate(c.ShouldBind(&request), c) {
		return
	}
	fmt.Println("request", request)
	result, isSucccess := SaveOnce(request)
	if !isSucccess {
		return
	}

	ResponseHandler.Go(c).SetData(&result).Apply()
}

func Update(c *gin.Context) {
	RequestJson.MustJSON = false

	request := UpdateItemRequest{}
	if RequestJson.Validate(c.ShouldBind(&request), c) {
		return
	}

	result, isSucccess := UpdateOnce(request)
	if !isSucccess {
		return
	}

	ResponseHandler.Go(c).SetData(&result).Apply()
}

func Delete(c *gin.Context) {
	RequestJson.MustJSON = false

	request := DetailItemRequest{}
	if RequestJson.Validate(c.ShouldBind(&request), c) {
		return
	}

	_, isSucccess := GetOnce(request)
	if !isSucccess {
		return
	}


	result, isSucccess := DeleteOnce(request)
	if !isSucccess {
		return
	}

	ResponseHandler.Go(c).SetData(&result).Apply()
}
