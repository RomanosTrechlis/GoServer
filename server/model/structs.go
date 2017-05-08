package helpers

import (
	"html/template"
	"os"
	"io/ioutil"
	"encoding/json"
)

var Config = Configuration{}
var ConfigFileName = "config.json"

type Page struct {
	Title       string
	Body        []byte //this means a byte slice
	DisplayBody template.HTML
}

type Post struct {
	Title   string
	Date    string
	Summary string
	Body    template.HTML
	File    string
}

type Blog struct {
	Blog template.HTML
}

type MarkdownPost struct {
	Title string
	Post  string
}

func (p *Page) Save() error {
	os.Mkdir(Config.Pages, 0777)
	filename := p.Title + ".txt"
	return ioutil.WriteFile(Config.Pages + filename, p.Body, 0600) // 0600 read write permissions
}

func (p *MarkdownPost) Save() error {
	os.Mkdir(Config.Posts, 0777)
	filename := p.Title + ".md"
	return ioutil.WriteFile(Config.Posts + filename, []byte(p.Post), 0600) // 0600 read write permissions
}

type Configuration struct {
	TextTemplates string `json:"TextTemplates`
	Templates     string `json:"Templates`
	Pages         string `json:"Pages`
	Posts         string `json:"Posts`
}

// ParseJSON unmarshals bytes to structs
func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

