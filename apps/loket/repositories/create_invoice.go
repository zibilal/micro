package repositories

import (
	"github.com/mataharimall/micro-api/services"
	"errors"
)

type ActionFunc func(inputs ...interface{}) (interface{}, error)

var CreateInvoice ActionFunc = func(inputs ...interface{}) (interface{}, error) {

	for _, inp := range inputs {
		invoice, found := inp.(services.LoketInvoice)

		if !found {
			return nil, errors.New("Wrong type arguments")
		}
	}

	return nil, nil

}