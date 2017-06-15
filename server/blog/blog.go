package blog

import (
	"net/http"

	"github.com/RomanosTrechlis/GoServer/server/util"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	util.Templates.ExecuteTemplate(w, "blog.html", buildBlog(r))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog/", http.StatusFound)
}
