package handlers

import (
	"net/http"

	"../logger"
	"../helpers"
)

var blogPath = "/blog/"

func NewBlogHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	p := helpers.CreateMarkdownPost()
	helpers.RenderTemplate(w, "newPost", p)
}

func SaveNewBlogHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	p := helpers.BuildMarkdownPost(r)
	err := p.Save()
	if err != nil {
		logger.Warning.Println("Error:", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, blogPath+p.Title, http.StatusFound)
}

func ReCacheHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	helpers.ReCacheTemplates()
	http.Redirect(w, r, blogPath, http.StatusFound)
}
