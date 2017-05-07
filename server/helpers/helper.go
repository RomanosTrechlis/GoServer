package helpers

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	txtTemplate "text/template"
	"time"
	"../logger"
)

var templatePath = "data/templates/"
var textTemplatePath = "data/textTemplates/"
var WikiValidPath = regexp.MustCompile(
	"^/wiki/(edit|save|view)/([a-zA-Z0-9]+)$")

var Templates = template.Must(template.ParseFiles(
	templatePath+"edit.html", templatePath+"view.html", templatePath+"blog.html"))
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
	err := Templates.ExecuteTemplate(w, tmpl+".html", p)
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
		title = m[len(m)-1]
	}
	return title
}

func ReCacheTemplates() {
	Templates = template.Must(template.ParseGlob("data/templates/*"))
	TextTemplates = txtTemplate.Must(txtTemplate.ParseGlob("data/textTemplates/*"))
}

func BuildMarkdownPost(r *http.Request) *MarkdownPost {
	post := r.FormValue("post")
	lines := strings.Split(post, "\r\n")
	title := strings.Replace(string(lines[0]), "Title: ", "", -1)
	title = strings.Replace(strings.ToLower(title), " ", "_", -1)
	return &MarkdownPost{Title: title, Post: post}
}

func CreateMarkdownPost() *MarkdownPost {
	var temp string
	buf := bytes.NewBufferString(temp)
	p := &MarkdownPost{Title: ""}
	TextTemplates.ExecuteTemplate(buf, "markdown.txt", p)
	t := time.Now()
	post := buf.String()
	post = strings.Replace(post, "Date:", "Date: "+t.String(), -1)
	return &MarkdownPost{Title: "", Post: post}
}
