package UserController

import (
	"inventori/app/Http/Controllers/User/UserModel"
)

func GetUserOnce(request UserModel.LoginRequest) (UserModel.User, error, bool) {
	getUser := UserModel.User{}
	errVal := UserModel.GormConnect.Table("users").Where("email = ?", request.Email).Find(&getUser)

	if errVal.Error != nil {
		return getUser, errVal.Error, false
	}

	return getUser, errVal.Error, true
}
