package helpers

import "html/template"

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
	Title 	string
	Post		string
}
