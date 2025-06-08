package api

import (
	"net/http"
	"sgl-rights/db"
	"strconv"
)

func RemoveUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		id, _ := strconv.Atoi(req.FormValue("id"))
		db.RemoveUser(id)
	}
}
