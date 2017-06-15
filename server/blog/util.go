package blog

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"

	"regexp"

	"github.com/RomanosTrechlis/GoServer/server"
	"github.com/RomanosTrechlis/GoServer/server/util"
)

type blog struct {
	Blog template.HTML
}

type post struct {
	Title   string
	Date    string
	Summary string
	Body    template.HTML
	File    string
}

var BlogValidPath = regexp.MustCompile(
	"^/blog/([a-zA-Z0-9_]+)$")

func getPosts(r *http.Request) []post {
	title := getPostName(r, BlogValidPath)
	a := []post{}
	fileName := title + ".md"
	if title == "" {
		fileName = "*"
	}

	files, _ := filepath.Glob(server.Config.Posts + fileName)
	for _, f := range files {
		a = getBlogPost(f, a)
	}
	return a
}

func getBlogPost(f string, a []post) []post {
	postsPath := strings.Replace(server.Config.Posts, "/", "\\", -1)
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
	a = append(a, post{title, date, summary, body, file})
	return a
}

func buildBlog(r *http.Request) *blog {
	posts := getPosts(r)
	var blogHtml string
	buf := bytes.NewBufferString(blogHtml)
	for i := 0; i < len(posts); i++ {
		util.TextTemplates.ExecuteTemplate(buf, "post.html", posts[i])
	}
	return &blog{Blog: template.HTML(buf.String())}
}

func getPostName(r *http.Request, regexp *regexp.Regexp) string {
	m := regexp.FindStringSubmatch(r.URL.Path)
	var title string
	if m == nil {
		title = ""
	} else {
		title = m[len(m)-1]
	}
	return title
}
