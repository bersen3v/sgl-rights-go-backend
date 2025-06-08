package api

import (
	"net/http"
	"os"
	"path/filepath"
)

func GetPhoto(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	filePath := filepath.Join("static", id)
	file, _ := os.Open(filePath)
	defer file.Close()
	fileInfo, _ := file.Stat()
	http.ServeContent(w, req, id, fileInfo.ModTime(), file)
}
