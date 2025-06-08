package api

import (
	"net/http"
	"sgl-rights/db"
	"strconv"
)

func RemoveEvent(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		id, _ := strconv.Atoi(req.FormValue("id"))
		db.RemoveEvent(id)
	}
}
