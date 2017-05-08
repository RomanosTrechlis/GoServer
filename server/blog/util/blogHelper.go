package helpers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"

	"regexp"
	"github.com/RomanosTrechlis/GoServer/server/util"
	structs "github.com/RomanosTrechlis/GoServer/server/model"
)

var BlogValidPath = regexp.MustCompile(
	"^/blog/([a-zA-Z0-9_]+)$")

func GetPosts(r *http.Request) []structs.Post {
	title := helpers.GetPostName(r, BlogValidPath)
	a := []structs.Post{}
	fileName := title + ".md"
	if title == "" {
		fileName = "*"
	}

	files, _ := filepath.Glob(structs.Config.Posts + fileName)
	for _, f := range files {
		a = GetBlogPost(f, a)
	}
	return a
}

func GetBlogPost(f string, a []structs.Post) []structs.Post {
	postsPath := strings.Replace(structs.Config.Posts, "/", "\\", -1)
	file := strings.Replace(f, postsPath, "", -1)
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
	a = append(a, structs.Post{title, date, summary, body, file})
	return a
}

func BuildBlog(r *http.Request) *structs.Blog {
	posts := GetPosts(r)
	var blogHtml string
	buf := bytes.NewBufferString(blogHtml)
	for i := 0; i < len(posts); i++ {
		helpers.TextTemplates.ExecuteTemplate(buf, "post.html", posts[i])
	}
	return &structs.Blog{Blog: template.HTML(buf.String())}
}
