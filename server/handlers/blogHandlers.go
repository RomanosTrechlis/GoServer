package handlers

import (
	"net/http"
	"../helpers"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	helpers.Templates.ExecuteTemplate(w, "blog.html", helpers.BuildBlog(r));
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog/", http.StatusFound)
}
