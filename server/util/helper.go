package helpers

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
	txtTemplate "text/template"

	"io/ioutil"

	"github.com/RomanosTrechlis/GoServer/server/logger"
	structs "github.com/RomanosTrechlis/GoServer/server/model"
)

var templatePath = "data/templates/"
var textTemplatePath = "data/textTemplates/"
var WikiValidPath = regexp.MustCompile(
	"^/wiki/(edit|save|view)/([a-zA-Z0-9]+)$")

var Templates = template.Must(template.ParseFiles(
	templatePath + "edit.html", templatePath + "view.html", templatePath + "blog.html"))
var TextTemplates = txtTemplate.Must(txtTemplate.ParseFiles(textTemplatePath + "post.html"))

var BlogValidPath = regexp.MustCompile(
	"^/blog/([a-zA-Z0-9_]+)$")
var adminNewBlogValidPath = regexp.MustCompile(
	"^/admin/blog/new/([a-zA-Z0-9_]+)$")
var adminSaveBlogValidPath = regexp.MustCompile(
	"^/admin/blog/save/([a-zA-Z0-9_]+)$")

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

func GetPostName(r *http.Request, regexp *regexp.Regexp) string {
	m := regexp.FindStringSubmatch(r.URL.Path)
	var title string
	if m == nil {
		title = ""
	} else {
		title = m[len(m) - 1]
	}
	return title
}

func LoadConfig(configPath string, loadTemplates bool) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Error.Println("Cannot load configuration file.")
		return
	}

	err = structs.Config.ParseJSON(data)
	if err != nil {
		logger.Error.Println("Failed to parse json file.")
	}
	if !loadTemplates {
		return
	}
	LoadTemplates()
}

func LoadTemplates() {
	Templates = template.Must(template.ParseGlob(structs.Config.Templates + "*"))
	TextTemplates = txtTemplate.Must(txtTemplate.ParseGlob(structs.Config.TextTemplates + "*"))
}

func LoadPage(title string) (*structs.Page, error) {
	filename := structs.Config.Pages + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &structs.Page{Title: title, Body: body}, nil
}

