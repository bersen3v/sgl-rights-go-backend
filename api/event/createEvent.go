package api

import (
	"net/http"
	"sgl-rights/db"
	"sgl-rights/entities"
	"strconv"
)

func CreateEvent(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		file, handler, _ := req.FormFile("photo")

		defer file.Close()
		previewPhoto := db.SavePhoto(handler, file)

		prize, _ := strconv.Atoi(req.FormValue("prize"))
		startTime, _ := strconv.Atoi(req.FormValue("startTime"))
		endTime, _ := strconv.Atoi(req.FormValue("endTime"))
		event := entities.Event{
			Id:           -1,
			PreviewPhoto: previewPhoto,
			Name: entities.I18nText{
				Ru: req.FormValue("nameRu"),
				En: req.FormValue("nameEn"),
				Kz: req.FormValue("nameKz"),
			},
			Description: entities.I18nText{
				Ru: req.FormValue("descriptionRu"),
				En: req.FormValue("descriptionEn"),
				Kz: req.FormValue("descriptionKz"),
			},
			Manager:   req.FormValue("manager"),
			Developer: req.FormValue("developer"),
			Place: entities.I18nText{
				Ru: req.FormValue("placeRu"),
				En: req.FormValue("placeEn"),
				Kz: req.FormValue("placeKz"),
			},
			Discipline: req.FormValue("discipline"),
			StartTime:  startTime,
			EndTime:    endTime,
			Prize:      prize,
		}

		db.AddEvent(event)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Событие успешно добавлено"))
		return
	}
}
