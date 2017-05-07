package handlers

import (
	"net/http"
	"../helpers"
)

var blogPath = "/blog/"

func NewBlogHandler(w http.ResponseWriter, r *http.Request) {
	p := helpers.CreateMarkdownPost()
	helpers.RenderTemplate(w, "newPost", p)
}

func SaveNewBlogHandler(w http.ResponseWriter, r *http.Request) {
	p := helpers.BuildMarkdownPost(r)
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, blogPath+p.Title, http.StatusFound)
}

func ReCacheHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ReCacheTemplates()
	http.Redirect(w, r, blogPath, http.StatusFound)
}
