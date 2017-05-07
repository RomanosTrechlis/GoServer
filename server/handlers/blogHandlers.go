package handlers

import (
	"net/http"

	"../../logger"
	"../helpers"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	helpers.Templates.ExecuteTemplate(w, "blog.html", helpers.BuildBlog(r));
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	http.Redirect(w, r, "/blog/", http.StatusFound)
}
