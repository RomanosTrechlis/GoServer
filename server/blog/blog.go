package blog

import (
	"net/http"


	"github.com/RomanosTrechlis/GoServer/server/logger"
	"github.com/RomanosTrechlis/GoServer/server/util"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Println(r.URL.Path)
	util.Templates.ExecuteTemplate(w, "blog.html", BuildBlog(r));
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Println(r.URL.Path)
	http.Redirect(w, r, "/blog/", http.StatusFound)
}
