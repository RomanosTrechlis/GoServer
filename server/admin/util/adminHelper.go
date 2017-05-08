package util

import (
	"net/http"
	"strings"
	"bytes"
	"time"

	"github.com/RomanosTrechlis/GoServer/server/util"
	structs "github.com/RomanosTrechlis/GoServer/server/model"
)

func BuildMarkdownPost(r *http.Request) *structs.MarkdownPost {
	post := r.FormValue("post")
	lines := strings.Split(post, "\r\n")
	title := strings.Replace(string(lines[0]), "Title: ", "", -1)
	title = strings.Replace(strings.ToLower(title), " ", "_", -1)
	return &structs.MarkdownPost{Title: title, Post: post}
}

func CreateMarkdownPost() *structs.MarkdownPost {
	var temp string
	buf := bytes.NewBufferString(temp)
	p := &structs.MarkdownPost{Title: ""}
	helpers.TextTemplates.ExecuteTemplate(buf, "markdown.txt", p)
	t := time.Now()
	post := buf.String()
	post = strings.Replace(post, "Date:", "Date: "+t.String(), -1)
	return &structs.MarkdownPost{Title: "", Post: post}
}
