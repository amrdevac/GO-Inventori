package ErrorHandler

import (
	"os"

	"github.com/withmandala/go-log"
)

type SetError struct {
	ErrorValue error
	checkName  string
}

func Err(argErrorValue error) SetError {
	var setERror SetError
	setERror.ErrorValue = argErrorValue
	return setERror
}

func (value SetError) Check(argCheckName string) SetError {
	value.checkName = argCheckName
	return value
}

func (value SetError) Fatal() {
	logger := log.New(os.Stderr)
	if value.ErrorValue != nil {
		logger.Fatal(value.ErrorValue)
	}
	logger.Info("Success Processing", value.checkName)
}

func (value SetError) Error() {
	logger := log.New(os.Stderr)
	logger.Info("Error When : ", value.checkName)
	if value.ErrorValue != nil {
		logger.Error(value.ErrorValue)
	}
}
