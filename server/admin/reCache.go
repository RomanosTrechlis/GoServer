package admin

import (
	"net/http"

	"github.com/RomanosTrechlis/GoServer/server"
	"github.com/RomanosTrechlis/GoServer/server/util"
)

func ReCacheHandler(w http.ResponseWriter, r *http.Request) {
	util.LoadConfig(server.ConfigFileName, false)
	http.Redirect(w, r, "/", http.StatusFound)
}

func LoadConfigHandler(w http.ResponseWriter, r *http.Request) {
	util.LoadTemplates()
	http.Redirect(w, r, "/", http.StatusFound)
}
