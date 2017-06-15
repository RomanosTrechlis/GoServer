package wiki

import (
	"html/template"
	"net/http"

	"github.com/RomanosTrechlis/GoServer/server/logger"
	"github.com/RomanosTrechlis/GoServer/server/util"
)

var viewPath = "/wiki/view/"
var editPath = "/wiki/edit/"
var savePath = "/wiki/save/"

//  handle URLs prefixed with "/view/"
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	// if page doesn't exists it should redirect to the edit page
	if err != nil {
		http.Redirect(w, r, editPath+title, http.StatusFound)
		return
	}

	p.DisplayBody = template.HTML(string(p.Body)) // make it display html
	util.RenderTemplate(w, "view", p)
}

//  handle URLs prefixed with "/edit/"
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	util.RenderTemplate(w, "edit", p)
}

//  handle URLs prefixed with "/save/"
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.Save()
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
		m := util.WikiValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			logger.Warning.Println("url not found")
			return
		}
		fn(w, r, m[2])
	})
}
