package AutentikasiService

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Http/Controllers/User/UserModel"
	"inventori/app/Provider/Hash"
	"inventori/app/Provider/ResponseHandler"
)

func MathcingPassword(request UserModel.LoginRequest) (UserModel.User, bool) {
	getUser := UserModel.User{}
	UserModel.GormConnect.Table("users").Where("email = ?", request.Email).Find(&getUser)

	isSucces, _ := Hash.Verify(request.Password, getUser.Password)
	if !isSucces {
		ResponseHandler.Go(controllers.GlobalGContext).CustomProccessFailure("Password didnt match")
		return getUser, false
	}
	return getUser, true
}

func EmailLoginVerify() {

}
