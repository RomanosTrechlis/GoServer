package blog

import (
	"net/http"

	c "github.com/RomanosTrechlis/GoServer/util/conf"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	c.Templates.ExecuteTemplate(w, "blog.html", buildBlog(r))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog/", http.StatusFound)
}
