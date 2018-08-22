package api

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/im-kulikov/helium/web"
)

func connectValidators(v web.Validator) error {
	if err := v.Register("address", validateAddress); err != nil {
		return err
	}

	// try to connect other..

	return nil
}

func validateAddress(fl web.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return common.IsHexAddress(field.String())
	default:
		return false
	}
}
