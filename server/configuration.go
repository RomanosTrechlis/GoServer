package server

import (
	"encoding/json"
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

