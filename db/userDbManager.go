package db

import (
	"database/sql"
	"fmt"
	"sgl-rights/entities"
)

func AddUser(u entities.User) {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	db.Exec(`
		INSERT INTO Users (previewPhoto, firstName, lastName, company, mail, phone, isAdmin, login, password) VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, u.PreviewPhoto, u.FirstName, u.LastName, u.Company, u.Mail, u.Phone, u.IsAdmin, u.Login, u.Password)
}

func GetAllUsers() []entities.User {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM Users")
	defer rows.Close()

	users := []entities.User{}

	for rows.Next() {
		e := entities.User{}
		rows.Scan(&e.Id, &e.PreviewPhoto, &e.FirstName, &e.LastName, &e.Company, &e.Mail, &e.Phone, &e.IsAdmin, &e.Login, &e.Password)
		users = append(users, e)
	}

	return users
}

func GetUserById(id int) entities.User {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Users WHERE id = $1", id)

	e := entities.User{}
	row.Scan(&e.Id, &e.PreviewPhoto, &e.FirstName, &e.LastName, &e.Company, &e.Mail, &e.Phone, &e.IsAdmin, &e.Login, &e.Password)

	return e
}

func UpdateUser(e entities.User) {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	_, i := db.Exec(`
		UPDATE Users
		SET previewPhoto = $1, 
			firstName = $2, 
			lastName = $3, 
			company = $4, 
			mail = $5, 
			phone = $6, 
			isAdmin = $7, 
			login = $8, 
			password = $9
		WHERE id = $10;
	`, e.PreviewPhoto, e.FirstName, e.LastName, e.Company, e.Mail, e.Phone, e.IsAdmin, e.Login, e.Password, e.Id)
	if i != nil {
		fmt.Println(i)
	}
}

func RemoveUser(id int) {
	db, _ := sql.Open("sqlite3", "store.db")

	defer db.Close()

	db.Exec("DELETE FROM Users WHERE id = $1", id)
}

func AuthUser(login string, password string) entities.User {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Users WHERE login = $1 AND password = $2", login, password)

	e := entities.User{}
	row.Scan(&e.Id, &e.PreviewPhoto, &e.FirstName, &e.LastName, &e.Company, &e.Mail, &e.Phone, &e.IsAdmin, &e.Login, &e.Password)

	return e
}

func AddEventToUser(userId int, eventId int, time int) {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	db.Exec(`
		INSERT INTO Sales (eventId, userId, date) VALUES
		($1, $2, $3)
	`, eventId, userId, time)
}

func RemoveEventFromUser(userId int, eventId int) {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	db.Exec(`
		DELETE FROM Sales 
        WHERE userId = $1 AND eventId = $2
	`, userId, eventId)
}

func GetUserEvents(userId int) []entities.Event {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	rows, _ := db.Query(`
		SELECT Events.id, previewPhoto, nameRu, nameEn, nameKz, descriptionRu, descriptionEn, descriptionKz, manager, developer, placeRu, placeEn, placeKz, discipline, startTime, endTime, prize FROM Sales
		JOIN Events ON Events.id = eventId
		WHERE userId = $1
	`, userId)
	defer rows.Close()

	events := []entities.Event{}

	for rows.Next() {
		e := entities.Event{}
		rows.Scan(&e.Id, &e.PreviewPhoto, &e.Name.Ru, &e.Name.En, &e.Name.Kz, &e.Description.Ru, &e.Description.En, &e.Description.Kz, &e.Manager, &e.Developer, &e.Place.Ru, &e.Place.En, &e.Place.Kz, &e.Discipline, &e.StartTime, &e.EndTime, &e.Prize)
		events = append(events, e)
	}

	return events
}
