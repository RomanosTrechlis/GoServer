package helpers

import (
	"io/ioutil"
	"os"
)

func (p *Page) Save() error {
	os.Mkdir("data/pages", 0777)
	filename := p.Title + ".txt"
	return ioutil.WriteFile("data/pages/"+filename, p.Body, 0600) // 0600 read write permissions
}

func (p *MarkdownPost) Save() error {
	os.Mkdir("data/posts", 0777)
	filename := p.Title + ".md"
	return ioutil.WriteFile("data/posts/"+filename, []byte(p.Post), 0600) // 0600 read write permissions
}

func LoadPage(title string) (*Page, error) {
	filename := "data/pages/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
