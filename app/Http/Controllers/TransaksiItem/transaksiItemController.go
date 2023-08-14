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

func DetailFull(c *gin.Context) {
	RequestJson.MustJSON = true
	request := DetailRequest{}
	if RequestJson.Validate(c.ShouldBindJSON(&request), c) {
		return
	}

	// get query Result
	listData, isSucccess := GetOnceWithDetailTransaksiItem(request)
	if !isSucccess {
		return
	}

	// Create a Template to store array value of 'detail_transaksi' (result of 1 to many query)
	var makeAnArrayOfObject []interface{}
	// Transform Struct into ObjctInterface{}
	mainParentData := helper.StructToInterfaceObj(listData, []string{})
	for _, vDetailTransaksi := range listData.DetailTransaksi {
		// Create a Template to store
		notNullFixedDetailItem := make(map[string]interface{})

		// Transform obj value of "detail transaksi"  as a interface
		arrObjDefaultDetailTransaksi := helper.StructToInterfaceObj(vDetailTransaksi, []string{})

		// Transform obj listData.DetailTransaksi" as a interface
		detailItem := helper.StructToInterfaceObj(vDetailTransaksi.DetailItem, []string{})
		if vDetailTransaksi.DetailItem.Nama != "" {
			notNullFixedDetailItem = detailItem
		}
		
		// Inject new fixed detail item that has value inside of it
		arrObjDefaultDetailTransaksi["detail_item"] = notNullFixedDetailItem
		// lastly , appen the object to make an array of "detail transaksi" that ben transformed
		makeAnArrayOfObject = append(makeAnArrayOfObject, arrObjDefaultDetailTransaksi)
	}
	mainParentData["detail_transaksi"] = makeAnArrayOfObject

	result := make(map[string]interface{})
	result["data"] = mainParentData
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
