package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sgl-rights/db"
	"strconv"
)

func GetEventById(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		id, _ := strconv.Atoi(req.FormValue("id"))
		event := db.GetEventById(id)
		jsonData, _ := json.Marshal(event)
		fmt.Fprintf(w, "%s", string(jsonData))
	}
}
