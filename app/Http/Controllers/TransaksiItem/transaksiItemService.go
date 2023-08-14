package transaksiitem

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Provider/ResponseHandler"
	"inventori/pkg/helper"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var privateGContext *gin.Context

func GetAll(request ListRequest) ([]TransaksiItemTable, bool) {
	listItem := []TransaksiItemTable{}
	gormResult := GormConnect.
		Limit(helper.PaginationLimit()).Offset(request.Offset * helper.PaginationLimit()).
		Order("created_at desc")

	if request.F_nama != "" {
		gormResult.Where("nama LIKE ?", "%"+request.F_nama+"%")
	}

	gormResult.Find(&listItem)
	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return listItem, false
	}

	return listItem, true
}

func GetOnce(request DetailRequest) (TransaksiItemTable, bool) {
	listItem := TransaksiItemTable{}
	gormResult := GormConnect.Where(TablePrimary, request.Id_Transaksi).First(&listItem)

	if gormResult.Error == gorm.ErrRecordNotFound {
		ResponseHandler.Go(privateGContext).DataNotFound(strconv.Itoa(request.Id_Transaksi))
		return listItem, false
	}

	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return listItem, false
	}

	return listItem, true
}

func GetOnceWithDetailTransaksiItem(request DetailRequest) (TransaksiItemWithDetail, bool) {
	listItem := TransaksiItemWithDetail{}
	gormResult := GormConnect.Preload("DetailTransaksi.DetailItem", "id_barang = ?", 4).Where(TablePrimary, request.Id_Transaksi).First(&listItem)

	if gormResult.Error == gorm.ErrRecordNotFound {
		ResponseHandler.Go(privateGContext).DataNotFound(strconv.Itoa(request.Id_Transaksi))
		return listItem, false
	}

	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return listItem, false
	}

	return listItem, true
}

func CountAll(request ListRequest) (int64, bool) {
	listItem := []TransaksiItemTable{}
	gormResult := GormConnect.Find(&listItem)

	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return 0, false
	}

	countAll := gormResult.RowsAffected

	return countAll, true
}

func InsertOnce(request InsertRequest) (map[string]interface{}, bool) {
	insertDB := helper.StructToInterfaceObj(request, []string{"RefBarang", "CreatedAt"})
	gormResultCreate := controllers.MaindGormConnect.Table(TableName).Create(&insertDB)

	if gormResultCreate.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResultCreate.Error)
		return insertDB, false
	}

	currentInsertData := map[string]interface{}{}
	gormResultGetCurrent := controllers.MaindGormConnect.Table(TableName).Order("created_at desc").Take(&currentInsertData)
	if gormResultGetCurrent.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResultGetCurrent.Error)
		return currentInsertData, false
	}

	return currentInsertData, true
}

func DeleteOnce(request DetailRequest) (DetailRequest, bool) {
	gormResult := GormConnect.Table(TableName).Where(TablePrimary, request.Id_Transaksi).Delete(&request)
	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return request, false
	}
	return request, true
}
