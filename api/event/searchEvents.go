package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sgl-rights/db"
	"strconv"
)

func SearchEvents(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		query := req.FormValue("query")
		disciplines := req.FormValue("disciplines")
		managers := req.FormValue("managers")
		developers := req.FormValue("developers")
		prizeMin, _ := strconv.Atoi(req.FormValue("prizeMin"))
		prizeMax, _ := strconv.Atoi(req.FormValue("prizeMax"))
		startTime, _ := strconv.Atoi(req.FormValue("startTime"))
		endTime, _ := strconv.Atoi(req.FormValue("endTime"))

		events := db.SearchEvents(query, disciplines, managers, developers, prizeMin, prizeMax, startTime, endTime)

		jsonData, _ := json.Marshal(events)
		fmt.Fprintf(w, "%s", string(jsonData))
	}
}
