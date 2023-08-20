package helper

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Provider/ResponseHandler"
)

func CompareChecker(argValue string, operator string, argUnexpectedCol string, argMessgae string) bool {
	var result bool
	switch operator {
	case "==":
		if argValue == argUnexpectedCol {
			ResponseHandler.Go(controllers.GlobalGContext).BadRequest(argMessgae)
			result = false
		} else {
			result = true
		}

	case "!=":
		if argValue != argUnexpectedCol {
			ResponseHandler.Go(controllers.GlobalGContext).BadRequest(argMessgae)
			result = false
		} else {
			result = true
		}
	case ">":
		if argValue > argUnexpectedCol {
			ResponseHandler.Go(controllers.GlobalGContext).BadRequest(argMessgae)
			result = false
		} else {
			result = true
		}
	case ">=":
		if argValue >= argUnexpectedCol {
			ResponseHandler.Go(controllers.GlobalGContext).BadRequest(argMessgae)
			result = false
		} else {
			result = true
		}
	case "<":
		if argValue < argUnexpectedCol {
			ResponseHandler.Go(controllers.GlobalGContext).BadRequest(argMessgae)

			result = false
		} else {
			result = true
		}
	case "<=":
		if argValue <= argUnexpectedCol {
			ResponseHandler.Go(controllers.GlobalGContext).BadRequest(argMessgae)
			result = false
		} else {
			result = true
		}
	}
	return result
}
