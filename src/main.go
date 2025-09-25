package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/alwaysnur/bookbank/src/upload"
	"github.com/alwaysnur/bookbank/web"
)

func main() {

	absPath, err := filepath.Abs("./store")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving from: ", absPath)

	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(absPath))))

	http.HandleFunc("/api/add", upload.UploadHandler)
	// static html routes
	http.HandleFunc("/new", web.HandleAdd)
	http.HandleFunc("/listen", web.HandleListenPage)
	http.HandleFunc("/", web.HandleIndex)
	http.HandleFunc("/library", web.HandleLibrary)
	// file servers
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
