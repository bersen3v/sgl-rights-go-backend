package api

import (
	"net/http"
	"sgl-rights/db"
	"strconv"
)

func RemoveEventFromUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		userId, _ := strconv.Atoi(req.FormValue("userId"))
		eventId, _ := strconv.Atoi(req.FormValue("eventId"))

		db.RemoveEventFromUser(userId, eventId)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Юзер успешно добавлен"))
	}
}
