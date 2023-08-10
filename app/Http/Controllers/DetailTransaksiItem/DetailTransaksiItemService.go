package detailtransaksiitem

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Provider/ResponseHandler"
	"inventori/pkg/helper"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var privateGContext *gin.Context

func GetAll(request ListRequest) ([]DetailTransaksiItemTable, bool) {
	listItem := []DetailTransaksiItemTable{}
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

func GetAllByIdTransaksi(request DetailRequest) ([]DetailTransaksiItemTable, bool) {
	listItem := []DetailTransaksiItemTable{}
	gormResult := GormConnect.Table(TableName).Where("id_transaksi", request.Id_Transaksi).First(&listItem)

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
	listItem := []DetailTransaksiItemTable{}
	gormResult := GormConnect.Find(&listItem)

	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return 0, false
	}

	countAll := gormResult.RowsAffected

	return countAll, true
}

func InsertOnceByIdTransaksi(requestRefBarang []InsertRequest, id_transaksi int) (int64, bool) {
	for index := range requestRefBarang {
		requestRefBarang[index].RefIdTransaksi = id_transaksi
	}

	gormResult := controllers.MaindGormConnect.Table(TableName).Create(&requestRefBarang)
	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return 0, false
	}
	return gormResult.RowsAffected, true
}

func DeleteOnce(request DetailRequest) (DetailRequest, bool) {
	gormResult := GormConnect.Table(TableName).Where(TablePrimary, request.Id_Transaksi).Delete(&request)
	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return request, false
	}
	return request, true
}
