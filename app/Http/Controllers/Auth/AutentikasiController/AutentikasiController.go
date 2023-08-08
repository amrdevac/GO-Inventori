package AutentikasiController

import (
	"inventori/app/Http/Controllers/Auth/AutentikasiService"
	"inventori/app/Http/Controllers/User/UserModel"
	"inventori/app/Provider/Hash"
	"inventori/app/Provider/RequestJson"
	"inventori/app/Provider/ResponseHandler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	users := UserModel.User{}
	RequestJson.Validate(c.ShouldBindJSON(&users), c)
	users.Password = Hash.Make(users.Password)
	UserModel.GormConnect.Create(&users)
	ResponseHandler.Go(c).SetData(&users).SetHttpStatus(http.StatusCreated).SetMessage("Successfully Registered").Apply()
}

func Login(c *gin.Context) {
	requestLogin := UserModel.LoginRequest{}
	c.ShouldBindJSON(&requestLogin)
	result, dataUser := AutentikasiService.MathcingPassword(requestLogin)

	if !result {
		ResponseHandler.Go(c).
			SetMessage("Username / Password tidak ditemukan").
			SetHttpStatus(http.StatusBadRequest).
			Apply().StopProcess()
	}

	LoginResponse := AutentikasiService.GetJWT(dataUser)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": LoginResponse, "message": "Login Berhasil"})
}
