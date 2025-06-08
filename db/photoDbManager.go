package db

import (
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
)

func SavePhoto(handler *multipart.FileHeader, file multipart.File) string {

	splittedName := strings.Split(handler.Filename, ".")
	format := splittedName[len(splittedName)-1]
	uniqueFileName := uuid.New().String()

	f, _ := os.OpenFile("static/"+uniqueFileName+"."+format, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	io.Copy(f, file)
	return uniqueFileName + "." + format
}
