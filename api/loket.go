package api

import (
	"encoding/json"
	"fmt"
	gr "github.com/parnurzeal/gorequest"
	c "github.com/spf13/viper"
	"net/http"
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

func New(configName string) *Loket {
	conf = c.GetStringMapString(configName)
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
