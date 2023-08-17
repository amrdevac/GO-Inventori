package AutentikasiService

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Http/Controllers/User/UserController"
	"inventori/app/Http/Controllers/User/UserModel"
	"inventori/app/Provider/Hash"
	"inventori/app/Provider/ResponseHandler"
	"inventori/pkg/helper"
)

func MathcingPassword(request UserModel.LoginRequest) (map[string]interface{}, bool) {
	getUser, errval, isSuccess := UserController.GetUserOnce(request)
	if !isSuccess {
		ResponseHandler.Go(controllers.GlobalGContext).DatabaseFetchFail(errval)
	}

	fixedUser := helper.StructToInterfaceObj(getUser, []string{"Password"})

	isSucces, _ := Hash.Verify(request.Password, getUser.Password)
	if !isSucces {
		ResponseHandler.Go(controllers.GlobalGContext).CustomProccessFailure("Password didnt match")
		return fixedUser, false
	}
	return fixedUser, true
}

func EmailLoginVerify() {

}
