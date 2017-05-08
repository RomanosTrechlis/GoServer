package admin

import (
	"net/http"

	"github.com/RomanosTrechlis/GoServer/server/logger"
	"github.com/RomanosTrechlis/GoServer/server/util"
	structs "github.com/RomanosTrechlis/GoServer/server/model"
)

func ReCacheHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("path:", r.URL.Path)
	helpers.LoadConfig(structs.ConfigFileName)
	http.Redirect(w, r, blogPath, http.StatusFound)
}