package isbn

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func GetCoverUrlByIsbn(isbn string) string {
	if isbn == "" {
		return "/static/image/placeholder.png"
	}
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%v", strings.ReplaceAll(isbn, "-", "")))
	if err != nil {
		log.Println("Error sending GET request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
	}
	coverImage := gjson.Get(string(body), "items.0.volumeInfo.imageLinks.thumbnail").String()
	if coverImage == "" {
		log.Println("An unknowen error occurred")
		return "/static/image/placeholder.png"
	}
	return coverImage
}

func GetDescriptionByIsbn(isbn string) string {
	if isbn == "" {
		return "No description provided (No ISBN provided)"
	}
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%v", strings.ReplaceAll(isbn, "-", "")))
	if err != nil {
		log.Println("Error sending GET request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
	}
	description := gjson.Get(string(body), "items.0.volumeInfo.description").String()
	if description == "" {
		log.Println("An unknowen error occurred")
		return "No description provided"
	}
	return description
}
