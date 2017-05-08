package admin

import (
	"net/http"

	"github.com/RomanosTrechlis/GoServer/server/logger"
	structs "github.com/RomanosTrechlis/GoServer/server/model"
	"github.com/RomanosTrechlis/GoServer/server/util"
)

func ReCacheHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	helpers.LoadConfig(structs.ConfigFileName, false)
	http.Redirect(w, r, "/", http.StatusFound)
}

func LoadConfigHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	helpers.LoadTemplates()
	http.Redirect(w, r, "/", http.StatusFound)
}