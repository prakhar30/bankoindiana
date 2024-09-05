package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/prakhar30/bankoindiana/utils"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsValidCurrency(currency)
	}
	return false
}
