package Hash

import (
	"inventori/app/Provider/ErrorHandler"

	"golang.org/x/crypto/bcrypt"
)

func Make(stringValue string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(stringValue), bcrypt.DefaultCost)
	ErrorHandler.Err(err).Check("Password Hash").Error()
	return string(pass)
}

func Verify(stringValue string, hashValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(stringValue))
	ErrorHandler.Err(err).Check("Password Didn't Match").Error()

	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword || err == bcrypt.ErrHashTooShort {
		return false
	}

	return true

}
