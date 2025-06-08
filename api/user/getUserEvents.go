package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sgl-rights/db"
	"strconv"
)

func GetUserEvents(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		userId, _ := strconv.Atoi(req.FormValue("userId"))
		events := db.GetUserEvents(userId)
		jsonData, _ := json.Marshal(events)
		fmt.Fprintf(w, "%s", string(jsonData))
	}
}
