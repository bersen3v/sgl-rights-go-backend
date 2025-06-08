package api

import (
	"net/http"
	"sgl-rights/db"
	"sgl-rights/entities"
	"strconv"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		file, handler, _ := req.FormFile("photo")

		defer file.Close()
		previewPhoto := db.SavePhoto(handler, file)

		isAdmin, _ := strconv.Atoi(req.FormValue("isAdmin"))
		user := entities.User{
			Id:           1,
			PreviewPhoto: previewPhoto,
			FirstName:    req.FormValue("firstName"),
			LastName:     req.FormValue("lastName"),
			Company:      req.FormValue("company"),
			Mail:         req.FormValue("mail"),
			Phone:        req.FormValue("phone"),
			Login:        req.FormValue("login"),
			Password:     req.FormValue("password"),
			IsAdmin:      isAdmin,
		}

		db.AddUser(user)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Юзер успешно добавлен"))
		return
	}
}
