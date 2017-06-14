package server

import (
	"encoding/json"
	"net/http"
)

var Config = Configuration{}
var ConfigFileName = "config.json"

type Configuration struct {
	TextTemplates string `json:"TextTemplates`
	Templates     string `json:"Templates`
	Pages         string `json:"Pages`
	Posts         string `json:"Posts`
}

// ParseJSON unmarshals bytes to structs
func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
