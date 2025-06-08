package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sgl-rights/db"
)

func GetAllEvents(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		events := db.GetAllEvents()
		jsonData, _ := json.Marshal(events)
		fmt.Fprintf(w, "%s", string(jsonData))
	}
}
