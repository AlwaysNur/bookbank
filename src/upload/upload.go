package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alwaysnur/bookbank/helper/books"
	"github.com/alwaysnur/bookbank/helper/log"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Error("Method not allowed")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<30)

	err := r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		log.Error("Error parsing form")
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		log.Error("Error retrieving the file")
		return
	}
	defer file.Close()

	name := r.FormValue("name")
	series := r.FormValue("series")
	author := r.FormValue("author")
	isbn := r.FormValue("isbn")
	// read counter file, convert to int
	counterPath := "./helper/counter"
	counterBytes, err := os.ReadFile(counterPath)
	var num int
	if err != nil {
		// if counter file is missing, start from 0
		num = 0
	} else {
		counterStr := strings.TrimSpace(string(counterBytes))
		num, err = strconv.Atoi(counterStr)
		if err != nil {
			http.Error(w, "Error converting counter to int", http.StatusInternalServerError)
			log.Fatal("Error converting counter to int")
			return
		}
	}

	nextNum := num + 1 // increment
	fileName := fmt.Sprint(nextNum) + ".mp3"
	dst, err := createFile(fileName)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		log.Error("Error creating counter file")
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		log.Error("Invalid or corrupted file")
		return
	}

	if !isValidFileType(fileBytes) && !strings.HasSuffix(header.Filename, ".mp3") {
		http.Error(w, "Invalid file type", http.StatusInternalServerError)
		log.Error(fmt.Sprintf("Invalid file type: '%s'", http.DetectContentType(fileBytes)))
		return
	}

	if _, err := dst.Write(fileBytes); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		log.Error("Error saving file")
		return
	}
	books.AddEntry(name, author, series, fileName, isbn) // add a new json entry
	err = os.WriteFile(counterPath, []byte(fmt.Sprint(nextNum)), 0644)
	if err != nil {
		log.Error("Failed to update counter")
		return
	}
	http.Redirect(w, r, "/listen?id="+fmt.Sprint(nextNum), http.StatusFound)
}

func createFile(filename string) (*os.File, error) {
	if _, err := os.Stat("store"); os.IsNotExist(err) {
		os.Mkdir("store", 0755)
	}
	dst, err := os.Create(filepath.Join("store", filename))
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func isValidFileType(file []byte) bool {
	fileType := http.DetectContentType(file)
	return fileType == "audio/mpeg"
}
