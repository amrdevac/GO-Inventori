package item

import (
	"inventori/app/Provider/ResponseHandler"
	"inventori/pkg/helper"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var privateGContext *gin.Context

func GetAll(request ListBarangRequest) ([]ItemTable, bool) {
	listItem := []ItemTable{}
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

func GetOnce(request DetailItemRequest) (ItemTable, bool) {
	listItem := ItemTable{}
	gormResult := GormConnect.Where("id_barang", request.IdBarang).First(&listItem)

	if gormResult.Error == gorm.ErrRecordNotFound {
		ResponseHandler.Go(privateGContext).DataNotFound(strconv.Itoa(request.IdBarang))
		return listItem, false
	}

	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return listItem, false
	}

	return listItem, true
}

func CountAll(request ListBarangRequest) (int64, bool) {
	listItem := []ItemTable{}
	gormResult := GormConnect.Find(&listItem)

	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return 0, false
	}

	countAll := gormResult.RowsAffected

	return countAll, true
}

func SaveOnce(request ItemTable) (ItemTable, bool) {

	request.CreatedAt = time.Now()
	gormResult := GormConnect.Create(&request)
	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return request, false
	}
	return request, true
}

func UpdateOnce(request UpdateItemRequest) (UpdateItemRequest, bool) {
	gormResult := GormConnect.Table("barang").Where("id_barang", request.IdBarang).Updates(&request)
	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return request, false
	}
	return request, true
}

func DeleteOnce(request DetailItemRequest) (DetailItemRequest, bool) {
	gormResult := GormConnect.Table("barang").Where("id_barang", request.IdBarang).Delete(&request)
	if gormResult.Error != nil {
		ResponseHandler.Go(privateGContext).DatabaseFetchFail(gormResult.Error)
		return request, false
	}
	return request, true
}
