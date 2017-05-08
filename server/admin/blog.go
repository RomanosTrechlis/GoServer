package admin

import (
	"net/http"


	"github.com/RomanosTrechlis/GoServer/server/logger"
	"github.com/RomanosTrechlis/GoServer/server/util"
	"github.com/RomanosTrechlis/GoServer/server/admin/util"
)

var blogPath = "/blog/"

func NewBlogHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	p := util.CreateMarkdownPost()
	helpers.RenderTemplate(w, "newPost", p)
}

func SaveNewBlogHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	p := util.BuildMarkdownPost(r)
	err := p.Save()
	if err != nil {
		logger.Warning.Println("Error:", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, blogPath+p.Title, http.StatusFound)
}


