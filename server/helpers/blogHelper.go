package helpers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

func GetPosts(r *http.Request) []Post {
	title := GetPostName(r, BlogValidPath)
	a := []Post{}
	fileName := title + ".md"
	if title == "" {
		fileName = "*"
	}
	files, _ := filepath.Glob("data/posts/" + fileName)
	for _, f := range files {
		a = GetBlogPost(f, a)
	}
	return a
}

func GetBlogPost(f string, a []Post) []Post {
	file := strings.Replace(f, "data\\posts\\", "", -1)
	file = strings.Replace(file, ".md", "", -1)
	fileRead, _ := ioutil.ReadFile(f)
	lines := strings.Split(string(fileRead), "\n")
	title := strings.Replace(string(lines[0]), "Title: ", "", -1)
	date := strings.Replace(string(lines[1]), "Date: ", "", -1)
	summary := strings.Replace(string(lines[2]), "Summary: ", "", -1)
	lineNumber := 3
	if string(lines[lineNumber]) == "---" {
		lineNumber = 4
	}
	bodyString := strings.Join(lines[lineNumber:len(lines)], "\n")
	body := template.HTML(blackfriday.MarkdownCommon([]byte(bodyString)))
	a = append(a, Post{title, date, summary, body, file})
	return a
}

func BuildBlog(r *http.Request) *Blog {
	posts := GetPosts(r)
	var blogHtml string
	buf := bytes.NewBufferString(blogHtml)
	for i := 0; i < len(posts); i++ {
		TextTemplates.ExecuteTemplate(buf, "post.html", posts[i])
	}
	return &Blog{Blog: template.HTML(buf.String())}
}
