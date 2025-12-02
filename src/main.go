package main

import (
	"net/http"
	"path/filepath"

	"github.com/alwaysnur/bookbank/helper/log"
	"github.com/alwaysnur/bookbank/src/upload"
	"github.com/alwaysnur/bookbank/web"
)

func main() {

	absPath, err := filepath.Abs("./store")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("Serving routes")
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(absPath))))
	// api routes
	http.HandleFunc("/api/add", upload.UploadHandler)
	http.HandleFunc("/api/version", web.HandleVersion)
	http.HandleFunc("/api/delete/", upload.DeleteHandler)
	// static html routes
	http.HandleFunc("/new", web.HandleAdd)
	http.HandleFunc("/listen", web.HandleListenPage)
	http.HandleFunc("/", web.HandleIndex)
	http.HandleFunc("/library", web.HandleLibrary)
	// file servers
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
