// Project's goal is to assist me in learning golang.
package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"./server/handlers"
	"./server/helpers"
	"./server/logger"
)

func main() {
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	// caching templates
	helpers.ReCacheTemplates()

	// routes
	// for a wiki we need three base routes view, edit, save
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/wiki/view/", handlers.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/wiki/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/wiki/save/", handlers.MakeHandler(handlers.SaveHandler))

	http.HandleFunc("/blog/", handlers.BlogHandler)
	http.HandleFunc("/admin/blog/new/", handlers.NewBlogHandler)
	http.HandleFunc("/admin/blog/save/", handlers.SaveNewBlogHandler)
	http.HandleFunc("/admin/recache/", handlers.ReCacheHandler)

	// allows css and js to be imported into pages
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	logger.Info.Println("Listening at port 8080...")
	http.ListenAndServe(":8080", nil)
}
