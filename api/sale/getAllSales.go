package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sgl-rights/db"
)

func GetAllSales(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		sales := db.GetAllSales()
		jsonData, _ := json.Marshal(sales)
		fmt.Fprintf(w, "%s", string(jsonData))
	}
}
