package api

import (
	"net/http"
	"sgl-rights/db"
	"sgl-rights/entities"
	"strconv"
)

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		previewPhoto := req.FormValue("previewPhoto")

		if previewPhoto == "" {
			file, handler, _ := req.FormFile("photo")
			defer file.Close()
			previewPhoto = db.SavePhoto(handler, file)
		}

		id, _ := strconv.Atoi(req.FormValue("id"))
		isAdmin, _ := strconv.Atoi(req.FormValue("isAdmin"))
		user := entities.User{
			Id:           id,
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

		db.UpdateUser(user)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Событие успешно изменено"))
		return
	}
}
