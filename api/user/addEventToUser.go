package api

import (
	"net/http"
	"sgl-rights/db"
	"strconv"
)

func AddEventToUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		userId, _ := strconv.Atoi(req.FormValue("userId"))
		eventId, _ := strconv.Atoi(req.FormValue("eventId"))
		time, _ := strconv.Atoi(req.FormValue("time"))

		db.AddEventToUser(userId, eventId, time)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Юзер успешно добавлен"))
	}
}
