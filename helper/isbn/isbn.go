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
		return "/static/image/placeholder.png"
	}
	return coverImage
}
