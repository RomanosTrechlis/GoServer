package wiki

import (
	"html/template"
	"io/ioutil"
	"os"

	"github.com/RomanosTrechlis/GoServer/server"
)

type Page struct {
	Title       string
	Body        []byte //this means a byte slice
	DisplayBody template.HTML
}

func LoadPage(title string) (*Page, error) {
	filename := server.Config.Pages + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func (p *Page) Save() error {
	os.Mkdir(server.Config.Pages, 0777)
	filename := p.Title + ".txt"
	return ioutil.WriteFile(server.Config.Pages+filename, p.Body, 0600) // 0600 read write permissions
}
