// Project's goal is to assist me in learning golang.
package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/RomanosTrechlis/GoServer/server/admin"
	"github.com/RomanosTrechlis/GoServer/server/blog"
	"github.com/RomanosTrechlis/GoServer/server/logger"
	structs "github.com/RomanosTrechlis/GoServer/server/model"
	"github.com/RomanosTrechlis/GoServer/server/util"
	"github.com/RomanosTrechlis/GoServer/server/wiki"
)

func initialize() {
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	// setting the configuration file
	structs.ConfigFileName = "config.json"
	// load configuration and chache templates
	helpers.LoadConfig(structs.ConfigFileName, true)
}

func main() {
	initialize()

	// routes
	// for a wiki we need three base routes view, edit, save
	http.HandleFunc("/", blog.RootHandler)
	http.HandleFunc("/wiki/view/", wiki.MakeHandler(wiki.ViewHandler))
	http.HandleFunc("/wiki/edit/", wiki.MakeHandler(wiki.EditHandler))
	http.HandleFunc("/wiki/save/", wiki.MakeHandler(wiki.SaveHandler))

	http.HandleFunc("/blog/", blog.BlogHandler)
	http.HandleFunc("/admin/blog/new/", admin.NewBlogHandler)
	http.HandleFunc("/admin/blog/save/", admin.SaveNewBlogHandler)
	http.HandleFunc("/admin/recache/", admin.ReCacheHandler)

	// allows css and js to be imported into pages
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	logger.Info.Println("Listening at port 8080...")
	http.ListenAndServe(":8080", nil)
}
