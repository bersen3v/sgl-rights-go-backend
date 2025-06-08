package api

import (
	"fmt"
	"net/http"
	"sgl-rights/db"
)

func GetEventsFilters(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		filters := db.GetFilters()
		fmt.Fprintf(w, "%s", string(filters))
	}
}
