package renderer

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/RomanosTrechlis/GoServer/logger"
	c "github.com/RomanosTrechlis/GoServer/util/conf"
)

var WikiValidPath = regexp.MustCompile(
	"^/wiki/(edit|save|view)/([a-zA-Z0-9]+)$")

// validates path
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := WikiValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := c.Templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		logger.Warning.Println("Error:", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
