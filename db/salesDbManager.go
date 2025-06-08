package db

import (
	"database/sql"
	"sgl-rights/entities"
)

func GetAllSales() []entities.Sale {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	rows, _ := db.Query(`
		SELECT * FROM Sales
	`)
	defer rows.Close()

	sales := []entities.Sale{}

	for rows.Next() {
		e := entities.Sale{}
		rows.Scan(&e.Id, &e.EventId, &e.UserId, &e.Time)
		sales = append(sales, e)
	}

	return sales
}
