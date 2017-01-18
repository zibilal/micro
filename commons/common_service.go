package commons

import "github.com/labstack/echo"

type Service interface {
	Request(data ...interface{}) (interface{}, error)
	Response(data interface{}) (interface{}, error)
}

type ServiceConfigurator interface {
	Configure(inputs map[string]interface{}) error
}

type AccessAction func(data ...interface{}) (interface{}, error)

type EchoControllerFunc func(c echo.Context) error
