package admin

import (
	"net/http"
	"strings"
	"bytes"
	"time"

	"github.com/RomanosTrechlis/GoServer/server/util"
	"github.com/RomanosTrechlis/GoServer/server"
	"os"
	"io/ioutil"
)

type MarkdownPost struct {
	Title string
	Post  string
}

func (p *MarkdownPost) Save() error {
	os.Mkdir(server.Config.Posts, 0777)
	filename := p.Title + ".md"
	return ioutil.WriteFile(server.Config.Posts + filename, []byte(p.Post), 0600) // 0600 read write permissions
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
	util.TextTemplates.ExecuteTemplate(buf, "markdown.txt", p)
	t := time.Now()
	post := buf.String()
	post = strings.Replace(post, "Date:", "Date: " + t.String(), -1)
	return &MarkdownPost{Title: "", Post: post}
}
