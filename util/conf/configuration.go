package conf

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	txtTemplate "text/template"

	"github.com/RomanosTrechlis/GoServer/logger"
)

var (
	Config           = configuration{}
	ConfigFileName   = "config.json"
	templatePath     = "data/templates/"
	textTemplatePath = "data/textTemplates/"
	Templates        = template.Must(template.ParseFiles(
		templatePath+"edit.html", templatePath+"view.html", templatePath+"blog.html"))
	TextTemplates = txtTemplate.Must(txtTemplate.ParseFiles(textTemplatePath + "post.html"))
)

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

func InitConfig(configPath string, loadTemplates bool) {
	if configPath != "" {
		LoadConfig(configPath, loadTemplates)
	}
	if !loadTemplates {
		return
	}
	LoadTemplates()
}

func LoadConfig(configPath string, loadTemplates bool) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Error.Println("Cannot load configuration file.")
		return
	}

	err = Config.ParseJSON(data)
	if err != nil {
		logger.Error.Println("Failed to parse json file.")
	}
}

func LoadTemplates() {
	Templates = template.Must(template.ParseGlob(Config.Templates + "*"))
	TextTemplates = txtTemplate.Must(txtTemplate.ParseGlob(Config.TextTemplates + "*"))
}
