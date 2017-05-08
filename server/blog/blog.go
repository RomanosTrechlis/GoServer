package blog

import (
	"net/http"


	"github.com/RomanosTrechlis/GoServer/server/logger"
	"github.com/RomanosTrechlis/GoServer/server/util"
	util "github.com/RomanosTrechlis/GoServer/server/blog/util"

)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	helpers.Templates.ExecuteTemplate(w, "blog.html", util.BuildBlog(r));
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	http.Redirect(w, r, "/blog/", http.StatusFound)
}
