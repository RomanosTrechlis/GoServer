package restricted

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"io/ioutil"
	"os"

	c "github.com/RomanosTrechlis/GoServer/util/conf"
	"golang.org/x/crypto/bcrypt"
)

type markdownPost struct {
	Title string
	Post  string
}

type user struct {
	Username string
	Password string
}

func (p *markdownPost) save() error {
	os.Mkdir(c.Config.Posts, 0777)
	filename := p.Title + ".md"
	return ioutil.WriteFile(c.Config.Posts+filename, []byte(p.Post), 0600) // 0600 read write permissions
}

func buildMarkdownPost(r *http.Request) *markdownPost {
	post := r.FormValue("post")
	lines := strings.Split(post, "\r\n")
	title := strings.Replace(string(lines[0]), "Title: ", "", -1)
	title = strings.Replace(strings.ToLower(title), " ", "_", -1)
	return &markdownPost{Title: title, Post: post}
}

func createMarkdownPost() *markdownPost {
	var temp string
	buf := bytes.NewBufferString(temp)
	p := &markdownPost{Title: ""}
	c.TextTemplates.ExecuteTemplate(buf, "markdown.txt", p)
	t := time.Now()
	post := buf.String()
	post = strings.Replace(post, "Date:", "Date: "+t.String(), -1)
	return &markdownPost{Title: "", Post: post}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
