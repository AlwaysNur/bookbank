package books

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	isbnlib "github.com/alwaysnur/bookbank/helper/isbn"
	"github.com/tidwall/gjson"
)

type JsonBook struct {
	Name     string `json:"name"`
	Author   string `json:"author"`
	Series   string `json:"series"`
	File     string `json:"file"`
	Isbn     string `json:"isbn"`
	CoverUrl string `json:"coverUrl"`
	Id       string `json:"id"`
}

type BooksData struct {
	Books []JsonBook `json:"books"`
}
type Book struct {
	Name     string
	Author   string
	Series   string
	File     string
	Isbn     string
	CoverUrl string
	Id       string
}

func check(err any) {
	if err != nil {
		panic(err)
	}
}

func AddEntry(name string, author string, series string, filename string, isbn string) {

	// create a new book entry
	newBook := JsonBook{
		Name:     name,
		Author:   author,
		Series:   series,
		File:     filename,
		Isbn:     isbn,
		CoverUrl: isbnlib.GetCoverUrlByIsbn(isbn),
		Id:       strings.ReplaceAll(filename, ".mp3", ""),
	}
	// read JSON file
	filePath := "helper/books.json"
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	var data BooksData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	// add the new book to the array
	data.Books = append(data.Books, newBook)

	// marshal back to JSON
	updatedJson, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}

	// write back to file
	if err := os.WriteFile(filePath, updatedJson, 0644); err != nil {
		panic(err)
	}

	log.Println("created book entry")
}

func GetBook(entry int) (string, string, string, string, string, string, string) {
	filePath := "helper/books.json"
	file, err := os.ReadFile(filePath)
	check(err)
	var (
		name     string = gjson.Get(string(file), fmt.Sprintf("books.%v.name", entry-1)).String()
		author   string = gjson.Get(string(file), fmt.Sprintf("books.%v.author", entry-1)).String()
		series   string = gjson.Get(string(file), fmt.Sprintf("books.%v.series", entry-1)).String()
		filepath string = gjson.Get(string(file), fmt.Sprintf("books.%v.file", entry-1)).String()
		isbn     string = gjson.Get(string(file), fmt.Sprintf("books.%v.isbn", entry-1)).String()
		coverUrl string = gjson.Get(string(file), fmt.Sprintf("books.%v.coverUrl", entry-1)).String()
		id       string = gjson.Get(string(file), fmt.Sprintf("books.%v.id", entry-1)).String()
	)
	return name, author, series, filepath, isbn, coverUrl, id
}

func GetBooks(filename string) ([]JsonBook, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data BooksData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data.Books, nil
}
