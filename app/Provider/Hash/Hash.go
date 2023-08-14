package Hash

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Provider/ResponseHandler"

	"golang.org/x/crypto/bcrypt"
)

func Make(stringValue string) (string, bool) {
	pass, err := bcrypt.GenerateFromPassword([]byte(stringValue), bcrypt.DefaultCost)
	if err != nil {
		ResponseHandler.Go(controllers.GlobalGContext).CustomProccessFailure(err.Error())
		return "", false
	}
	return string(pass), true
}

func Verify(stringValue string, hashValue string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(stringValue))

	if err != nil {
		return false, err.Error()
	}
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword || err == bcrypt.ErrHashTooShort {
		return false, "Password didnt match"
	}

	return true, "Password match"

}
