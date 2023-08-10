package transaksiitem

import (
	controllers "inventori/app/Http/Controllers"
	detailtransaksiitem "inventori/app/Http/Controllers/DetailTransaksiItem"
	"inventori/app/Provider/RequestJson"
	"inventori/app/Provider/ResponseHandler"
	"inventori/pkg/helper"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	RequestJson.MustJSON = true
	request := ListRequest{}
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
	request := DetailRequest{}
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
	RequestJson.MustJSON = false
	request := InsertRequest{}
	if RequestJson.Validate(c.ShouldBind(&request), c) {
		return
	}

	if helper.InArray(ArrTipeTransaksi, request.Tipe_Transaksi) < 0 {
		ResponseHandler.Go(c).EnumNotFound("tipe_transaksi", ArrTipeTransaksi)
		return
	}

	controllers.MaindGormConnect = GormConnect.Begin()
	result, isSucccess := InsertOnce(request)
	if !isSucccess {
		return
	}

	idTransaksi := result["id_transaksi"].(int32)
	_, isSucccess = detailtransaksiitem.InsertOnceByIdTransaksi(request.RefBarang, int(idTransaksi))
	if !isSucccess {
		return
	}
	controllers.MaindGormConnect.Commit()

	ResponseHandler.Go(c).SetData(&result).Apply()
}

func Delete(c *gin.Context) {
	RequestJson.MustJSON = false

	request := DetailRequest{}
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
