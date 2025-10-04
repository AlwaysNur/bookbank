package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/alwaysnur/bookbank/helper/books"
)

type listenData struct {
	Name     string
	Author   string
	Series   string
	File     string
	Isbn     string
	CoverUrl string
	Id       string
}
type indexData struct {
	Name              string
	ContinueListening any
}

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/add.html")
}
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
	}

	data := indexData{
		Name:              "Nur", // [TODO] Use Config.toml file for name
		ContinueListening: "hii", // [TODO] Use some method to get this data
	}
	tmpl.Execute(w, data)
}

func HandleLibrary(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/library.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
	}

	booksArray, err := books.GetBooks("helper/books.json")
	if err != nil {
		log.Printf("Error in json file: %v", err)
		return
	}

	if err := tmpl.Execute(w, booksArray); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func HandleListenPage(w http.ResponseWriter, r *http.Request) {

	u := r.URL

	queryParams := u.Query()

	id := queryParams.Get("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error converting counter to int", http.StatusInternalServerError)
		return
	}
	name, author, series, filename, isbn, coverUrl, _ := books.GetBook(num)

	tmpl, err := template.ParseFiles("web/listen.html")
	if err != nil {
		log.Println("Error parsing html")
		return
	}
	filepath := fmt.Sprintf("/file/%v", filename)
	data := listenData{
		Name:     name,
		Author:   author,
		Series:   series,
		File:     filepath,
		Isbn:     isbn,
		CoverUrl: coverUrl,
		Id:       strings.ReplaceAll(filename, ".mp3", ""),
	}
	tmpl.Execute(w, data)
}
