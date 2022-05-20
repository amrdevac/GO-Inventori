package AutentikasiService

import (
	"login-sistem-jwt/app/Http/Controllers/User/UserModel"
	"login-sistem-jwt/app/Provider/Hash"
)

func MathcingPassword(request UserModel.LoginRequest) (bool, UserModel.User) {
	getUser := UserModel.User{}
	UserModel.GormConnect.Table("users").Where("email = ?", request.Email).Find(&getUser)

	return Hash.Verify(request.Password, getUser.Password), getUser
}

func EmailLoginVerify() {

}
