package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb() {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	db.Exec(`
		CREATE TABLE IF NOT EXISTS Events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			previewPhoto TEXT NOT NULL,
			nameRu TEXT NOT NULL,
			nameEn TEXT NOT NULL,
			nameKz TEXT NOT NULL,
			descriptionRu TEXT NOT NULL,
			descriptionEn TEXT NOT NULL,
			descriptionKz TEXT NOT NULL,
			manager TEXT NOT NULL,
			developer TEXT NOT NULL,
			placeRu TEXT NOT NULL,
			placeEn TEXT NOT NULL,
			placeKz TEXT NOT NULL,
			discipline TEXT NOT NULL,
			startTime INTEGER NOT NULL,
			endTime INTEGER NOT NULL,
			prize INTEGER NOT NULL
		)	
	`)

	db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			previewPhoto TEXT NOT NULL,
			firstName TEXT NOT NULL,
			lastName TEXT NOT NULL,
			company TEXT NOT NULL,
			mail TEXT NOT NULL,
			phone TEXT NOT NULL,
			isAdmin INTEGER NOT NULL,
            login TEXT NOT NULL,
            password TEXT NOT NULL
		)
	`)

	db.Exec(`
		CREATE TABLE IF NOT EXISTS Sales (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			eventId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			date INTEGER NOT NULL,
			CONSTRAINT fk_eventId
				FOREIGN KEY (eventId)
				REFERENCES Events(id),
			CONSTRAINT fk_eventId
				FOREIGN KEY (userId)
				REFERENCES Users(id)
		)
	`)

}
