package restricted

import (
	"net/http"

	c "github.com/RomanosTrechlis/GoServer/util/conf"
)

func ReCacheHandler(w http.ResponseWriter, r *http.Request) {
	c.LoadConfig(c.ConfigFileName, false)
	http.Redirect(w, r, "/", http.StatusFound)
}

func LoadConfigHandler(w http.ResponseWriter, r *http.Request) {
	c.LoadTemplates()
	http.Redirect(w, r, "/", http.StatusFound)
}
