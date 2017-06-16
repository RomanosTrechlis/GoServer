package wiki

import (
	"html/template"
	"net/http"

	"github.com/RomanosTrechlis/GoServer/logger"
	rend "github.com/RomanosTrechlis/GoServer/util/renderer"
)

var viewPath = "/wiki/view/"
var editPath = "/wiki/edit/"
var savePath = "/wiki/save/"

//  handle URLs prefixed with "/view/"
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	// if page doesn't exists it should redirect to the edit page
	if err != nil {
		http.Redirect(w, r, editPath+title, http.StatusFound)
		return
	}

	p.DisplayBody = template.HTML(string(p.Body)) // make it display html
	rend.RenderTemplate(w, "view", p)
}

//  handle URLs prefixed with "/edit/"
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &page{Title: title}
	}
	rend.RenderTemplate(w, "edit", p)
}

//  handle URLs prefixed with "/save/"
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		logger.Warning.Println("Error:", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, viewPath+title, http.StatusFound)
}

// wrapper function
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug.Println(r.URL.Path)
		m := rend.WikiValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			logger.Warning.Println("url not found")
			return
		}
		fn(w, r, m[2])
	})
}
