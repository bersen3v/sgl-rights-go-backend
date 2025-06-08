package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sgl-rights/db"
)

func AuthUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		login := req.FormValue("login")
		password := req.FormValue("password")

		user := db.AuthUser(login, password)
		jsonData, _ := json.Marshal(user)
		fmt.Fprintf(w, "%s", string(jsonData))
	}
}
