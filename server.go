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
	"github.com/RomanosTrechlis/GoServer/server"

	"github.com/boltdb/bolt"

	"time"
	"log"
	"errors"
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
	structs.ConfigFileName = "config.json"
	// load configuration and chache templates
	helpers.LoadConfig(structs.ConfigFileName, true)

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

	mx := http.NewServeMux()
	// routes
	// for a wiki we need three base routes view, edit, save
	mx.HandleFunc("/", blog.RootHandler)
	mx.Handle("/wiki/view/", server.Adapt(wiki.MakeHandler(wiki.ViewHandler), wiki.WikiAdapter()))
	//mx.Handle("/wiki/view/", wiki.WikiAdapter()(wiki.MakeHandler(wiki.ViewHandler)))
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
