package commons

type Service interface {
	Request() (interface{}, error)
	Response() (interface{}, error)
}

type ServiceConfigurator interface {
	Configure(inputs map[string]interface{}) error
}

type AccessAction func(data ...interface{}) (Service, error)
