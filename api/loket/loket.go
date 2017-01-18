package loket

import (
	"encoding/json"
	"fmt"
	gr "github.com/parnurzeal/gorequest"
	c "github.com/spf13/viper"
	"net/http"
	"reflect"
	"github.com/mataharimall/micro-api/helpers"
)

var conf map[string]string

type Loket struct {
	UserName     string
	Password     string
	ApiKey       string
	Token        string
	Response     *http.Response
	Body         string
	Errors       []error
	TokenExpired bool
}

func (c *Loket) Configure(inputs map[string]interface{}) *Loket {
	fmt.Println("Inputs", inputs)
	if len(inputs) != 3 {
		return nil
	}

	val := reflect.Indirect(reflect.ValueOf(c))
	t := val.Type()

	if val.Kind() == reflect.Struct {
		for i:=0;i < t.NumField(); i++ {
			field := t.Field(i)
			var dest reflect.Value

			fName := helpers.FieldName("app", field)

			switch field.Type.Kind() {
			case reflect.String:
				if inputs[fName] != nil {
					tmp := inputs[fName].(string)
					dest = reflect.ValueOf(tmp)
					val.Field(i).Set(dest)
				}
			case reflect.Uint:
				if inputs[fName] != nil {
					tmp := inputs[fName].(uint)
					dest = reflect.ValueOf(tmp)
					val.Field(i).Set(dest)
				}
			}

		}
	}

	return c
}

func getConfig(key string) string {
	if _, ok := conf[key]; ok {
		return conf[key]
	}
	return ""
}

func (l *Loket) GetAuth() *Loket {
	if l.TokenExpired {
		return l
	}
	if len(l.UserName) == 0 || len(l.Password) == 0 || len(l.ApiKey) == 0 {
		return l
	}
	body := fmt.Sprintf(`{"username": "%s","password": "%s","APIKEY": "%s"}`, l.UserName, l.Password, l.ApiKey)
	l.Post("/v3/login", "form", body)
	l.SetToken()
	return l
}

func New() *Loket {
	conf = c.GetStringMapString("loket")
	l := &Loket{
		UserName: getConfig("username"),
		Password: getConfig("password"),
		ApiKey:   getConfig("key"),
		Token:    "",
	}
	return l
}

func SetUrl(url string) string {
	t := fmt.Sprintf("%s%s", getConfig("url"), url)
	return t
}

func (l *Loket) SetToken() *Loket {
	resp := struct {
		Status string `json:"status"`
		Data   *struct {
			Token string `json:"token"`
		} `json:"data"`
		Message string `json:"message"`
	}{"", nil, ""}
	byt := []byte(l.Body)

	if err := json.Unmarshal(byt, &resp); err != nil {
		return l
	}

	l.Token = resp.Data.Token
	return l
}

func (l *Loket) SetStruct(v interface{}) *Loket {
	err := json.Unmarshal([]byte(l.Body), &v)
	if err != nil {
		l.Errors = append(l.Errors, err)
		return l
	}
	return l
}

func (l *Loket) Post(url, t, body string) *Loket {
	l.Response, l.Body, l.Errors = gr.New().
		Post(SetUrl(url)).
		Type(t).
		Send(body).
		End()
	return l
}

func (l *Loket) Get(url string) *Loket {
	l.Response, l.Body, l.Errors = gr.New().
		Get(SetUrl(url)).
		End()
	return l
}
