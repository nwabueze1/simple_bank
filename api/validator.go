package api

import (
	"fidelis.com/simple_bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool{
	currrency, ok :=fieldLevel.Field().Interface().(string)

	if ok{
		//check currency is supported
		return util.IsSupportedCurrency(currrency)
	}
	return false
}