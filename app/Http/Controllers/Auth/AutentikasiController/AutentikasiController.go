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
	RequestJson.MustJSON = true
	users := UserModel.User{}
	if RequestJson.Validate(c.ShouldBindJSON(&users), c) {
		return
	}

	hashPass, isSuccess := Hash.Make(users.Password)
	if !isSuccess {
		return
	}

	users.Password = hashPass
	UserModel.GormConnect.Create(&users)
	ResponseHandler.Go(c).SetData(&users).SetHttpStatus(http.StatusCreated).SetMessage("Successfully Registered").Apply()
}

func Login(c *gin.Context) {
	requestLogin := UserModel.LoginRequest{}
	if RequestJson.Validate(c.ShouldBindJSON(&requestLogin), c) {
		return
	}

	dataUser, isSuccess := AutentikasiService.MathcingPassword(requestLogin)
	if !isSuccess {
		return
	}

	LoginResponse, isSuccess := AutentikasiService.GetJWT(dataUser)
	if !isSuccess {
		return
	}
	ResponseHandler.Go(c).SetData(&LoginResponse).Apply()

	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": LoginResponse, "message": "Login Berhasil"})
}
