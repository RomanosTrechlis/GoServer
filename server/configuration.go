package server

import (
	"encoding/json"
	"net/http"
)

var Config = configuration{}
var ConfigFileName = "config.json"

type configuration struct {
	TextTemplates string `json:"TextTemplates`
	Templates     string `json:"Templates`
	Pages         string `json:"Pages`
	Posts         string `json:"Posts`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
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
