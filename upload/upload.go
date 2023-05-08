package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the file has the ".so" extension
	if filepath.Ext(header.Filename) != ".so" {
		http.Error(w, "Only .so files are allowed", http.StatusUnsupportedMediaType)
		return
	}

	// Save the file
	dst, err := os.Create(fmt.Sprintf("sofiles/%s", header.Filename))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", header.Filename)
}
