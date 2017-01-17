package loket

import (
	"github.com/mataharimall/micro-api/commons"
	"errors"
)

type LoketEndpoints struct {
	Server commons.ServiceConfigurator
}

type LoketConfiguration struct {
	BaseUrl  string `app:"base_url"`
	ApiKey   string `app:"api_key"`
	UserName string `app:"user_name"`
	Password string `app:"password"`
}

func (c *LoketConfiguration) Configure(inputs map[string]interface{}) error {
	if len(inputs) != 4 {
		return errors.New("Please provide only base_url|api_key|user_name|password only")
	}

	return nil
}

type Loket struct {

}
