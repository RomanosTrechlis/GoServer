package util

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
	txtTemplate "text/template"

	"io/ioutil"

	"github.com/RomanosTrechlis/GoServer/server/logger"
	"github.com/RomanosTrechlis/GoServer/server"
)

var templatePath = "data/templates/"
var textTemplatePath = "data/textTemplates/"
var WikiValidPath = regexp.MustCompile(
	"^/wiki/(edit|save|view)/([a-zA-Z0-9]+)$")

var Templates = template.Must(template.ParseFiles(
	templatePath + "edit.html", templatePath + "view.html", templatePath + "blog.html"))
var TextTemplates = txtTemplate.Must(txtTemplate.ParseFiles(textTemplatePath + "post.html"))

// validates path
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := WikiValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := Templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		logger.Warning.Println("Error:", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoadConfig(configPath string, loadTemplates bool) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Error.Println("Cannot load configuration file.")
		return
	}

	err = server.Config.ParseJSON(data)
	if err != nil {
		logger.Error.Println("Failed to parse json file.")
	}
	if !loadTemplates {
		return
	}
	LoadTemplates()
}

func LoadTemplates() {
	Templates = template.Must(template.ParseGlob(server.Config.Templates + "*"))
	TextTemplates = txtTemplate.Must(txtTemplate.ParseGlob(server.Config.TextTemplates + "*"))
}
