package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sgl-rights/db"
)

func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		users := db.GetAllUsers()
		jsonData, _ := json.Marshal(users)
		fmt.Fprintf(w, "%s", string(jsonData))
	}
}
