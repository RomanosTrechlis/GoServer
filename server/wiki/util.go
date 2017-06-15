package wiki

import (
	"html/template"
	"io/ioutil"
	"os"

	"github.com/RomanosTrechlis/GoServer/server"
)

type page struct {
	Title       string
	Body        []byte //this means a byte slice
	DisplayBody template.HTML
}

func loadPage(title string) (*page, error) {
	filename := server.Config.Pages + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}

func (p *page) save() error {
	os.Mkdir(server.Config.Pages, 0777)
	filename := p.Title + ".txt"
	return ioutil.WriteFile(server.Config.Pages+filename, p.Body, 0600) // 0600 read write permissions
}
