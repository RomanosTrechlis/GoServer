// Project's goal is to assist me in learning golang.
package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/RomanosTrechlis/GoServer/server"
	"github.com/RomanosTrechlis/GoServer/server/admin"
	"github.com/RomanosTrechlis/GoServer/server/blog"
	"github.com/RomanosTrechlis/GoServer/server/logger"
	"github.com/RomanosTrechlis/GoServer/server/util"
	"github.com/RomanosTrechlis/GoServer/server/wiki"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"

	"errors"
	"log"
	"time"
)

func initialize() {
	/*path := os.Getwd() + "\\logs\\"
	fileName := "temp.log"
	file, err := os.OpenFile(path + fileName, os.O_RDONLY | os.O_CREATE, 0666)*/
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	/*if err != nil {
		logger.Warning.Println(err.Error())
	}*/

	// setting the configuration file
	server.ConfigFileName = "config.json"
	// load configuration and chache templates
	util.LoadConfig(server.ConfigFileName, true)

	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		logger.Warning.Println(err)
	}
	defer db.Close()
	world := []byte("world")
	key := []byte("hello")
	value := []byte("Hello World!")
	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logger.Warning.Println(err)
	}
	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			logger.Warning.Println("Bucket %q not found!", world)
			return errors.New("Error")
		}

		val := bucket.Get(key)
		logger.Info.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	initialize()

	mx := mux.NewRouter()
	// routes
	// for a wiki we need three base routes view, edit, save
	mx.HandleFunc("/", blog.RootHandler)
	mx.HandleFunc("/wiki/view/{^[a-zA-Z0-9_.-]*$}", server.Chain(wiki.MakeHandler(wiki.ViewHandler), admin.Logging())) // test
	//mx.Handle("/wiki/view/", wiki.WikiAdapter()(wiki.MakeHandler(wiki.ViewHandler)))
	mx.HandleFunc("/wiki/edit/{^[a-zA-Z0-9_.-]*$}", server.Chain(wiki.MakeHandler(wiki.EditHandler), admin.Logging()))
	mx.HandleFunc("/wiki/save/{^[a-zA-Z0-9_.-]*$}", server.Chain(wiki.MakeHandler(wiki.SaveHandler), admin.Logging()))

	mx.HandleFunc("/blog/", server.Chain(blog.BlogHandler, admin.Logging()))
	mx.HandleFunc("/blog/{^[a-zA-Z0-9_.-]*$}", server.Chain(blog.BlogHandler, admin.Logging()))

	mx.HandleFunc("/admin/blog/new/", server.Chain(admin.NewBlogHandler, admin.Authenticate(), admin.Logging()))
	mx.HandleFunc("/admin/blog/save/", server.Chain(admin.SaveNewBlogHandler, admin.Authenticate(), admin.Logging()))
	mx.HandleFunc("/admin/recache/", server.Chain(admin.ReCacheHandler, admin.Authenticate(), admin.Logging()))

	mx.HandleFunc("/login/", server.Chain(admin.LoginGet, admin.Logging())).Methods("GET")
	mx.HandleFunc("/login/", server.Chain(admin.LoginPost, admin.Logging())).Methods("POST")
	mx.HandleFunc("/logout/", server.Chain(admin.Logout, admin.Logging()))

	// allows css and js to be imported into pages
	mx.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	logger.Info.Println("Listening at port 8080...")
	http.ListenAndServe(":8080", mx)
}
