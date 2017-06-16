package restricted

import (
	"net/http"

	"github.com/RomanosTrechlis/GoServer/logger"
	rend "github.com/RomanosTrechlis/GoServer/util/renderer"
)

var blogPath = "/blog/"

func NewBlogHandler(w http.ResponseWriter, r *http.Request) {
	p := createMarkdownPost()
	rend.RenderTemplate(w, "newPost", p)
}

func SaveNewBlogHandler(w http.ResponseWriter, r *http.Request) {
	p := buildMarkdownPost(r)
	err := p.save()
	if err != nil {
		logger.Warning.Println("Error:", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, blogPath+p.Title, http.StatusFound)
}
