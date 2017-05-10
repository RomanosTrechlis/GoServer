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
	path := os.Getwd() + "\\logs\\"
	fileName := "temp.log"
	file, err := os.OpenFile(path + fileName, os.O_RDONLY | os.O_CREATE, 0666)
	logger.Init(ioutil.Discard, file, file, os.Stdout, os.Stderr)
	if err != nil {
		logger.Warning.Println(err.Error())
	}

	// setting the configuration file
	structs.ConfigFileName = "config.json"
	// load configuration and chache templates
	helpers.LoadConfig(structs.ConfigFileName, true)
}

func main() {
	initialize()

	mx := http.NewServeMux()
	// routes
	// for a wiki we need three base routes view, edit, save
	mx.HandleFunc("/", blog.RootHandler)
	mx.Handle("/wiki/view/", wiki.MakeHandler(wiki.ViewHandler))
	mx.Handle("/wiki/edit/", wiki.MakeHandler(wiki.EditHandler))
	mx.Handle("/wiki/save/", wiki.MakeHandler(wiki.SaveHandler))

	mx.HandleFunc("/blog/", blog.BlogHandler)
	mx.HandleFunc("/admin/blog/new/", admin.NewBlogHandler)
	mx.HandleFunc("/admin/blog/save/", admin.SaveNewBlogHandler)
	mx.HandleFunc("/admin/recache/", admin.ReCacheHandler)

	// allows css and js to be imported into pages
	mx.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	logger.Info.Println("Listening at port 8080...")
	http.ListenAndServe(":8080", mx)
}
