package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sgl-rights/entities"
	"strings"
)

func AddEvent(e entities.Event) {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	db.Exec(`
		INSERT INTO Events (
    		previewPhoto, nameRu, nameEn, nameKz, descriptionRu, descriptionEn, descriptionKz, manager, developer, placeRu, placeEn, placeKz, discipline, startTime, endTime, prize) VALUES
		(
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7, 
			$8, 
			$9, 
			$10, 
			$11, 
			$12, 
			$13, 
			$14,
			$15,
			$16
		)
	`, e.PreviewPhoto, e.Name.Ru, e.Name.En, e.Name.Kz, e.Description.Ru, e.Description.En, e.Description.Kz, e.Manager, e.Developer, e.Place.Ru, e.Place.En, e.Place.Kz, e.Discipline, e.StartTime, e.EndTime, e.Prize)
}

func UpdateEvent(e entities.Event) {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	_, i := db.Exec(`
		UPDATE Events
		SET previewPhoto = $1, 
			nameRu = $2, 
			nameEn = $3, 
			nameKz = $4, 
			descriptionRu = $5, 
			descriptionEn = $6, 
			descriptionKz = $7, 
			manager = $8, 
			developer = $9,
			placeRu = $10, 
			placeEn = $11, 
			placeKz = $12, 
			discipline = $13, 
			startTime = $14, 
			endTime = $15, 
			prize = $16 
		WHERE id = $17;
	`, e.PreviewPhoto, e.Name.Ru, e.Name.En, e.Name.Kz, e.Description.Ru, e.Description.En, e.Description.Kz, e.Manager, e.Developer, e.Place.Ru, e.Place.En, e.Place.Kz, e.Discipline, e.StartTime, e.EndTime, e.Prize, e.Id)
	if i != nil {
		fmt.Println(i)
	}
}

func GetAllEvents() []entities.Event {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM Events")
	defer rows.Close()

	events := []entities.Event{}

	for rows.Next() {
		e := entities.Event{}
		rows.Scan(&e.Id, &e.PreviewPhoto, &e.Name.Ru, &e.Name.En, &e.Name.Kz, &e.Description.Ru, &e.Description.En, &e.Description.Kz, &e.Manager, &e.Developer, &e.Place.Ru, &e.Place.En, &e.Place.Kz, &e.Discipline, &e.StartTime, &e.EndTime, &e.Prize)
		events = append(events, e)
	}

	return events
}

func SearchEvents(query string, disciplines string, managers string, developers string, prizeMin int, prizeMax int, startTime int, endTime int) []entities.Event {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	fmt.Println(disciplines)

	disciplinesFilter := ""
	if disciplines != "" {
		disciplinesFilter = fmt.Sprintf("discipline IN (%s)", disciplines)
	}

	managersFilter := ""
	if managers != "" {
		managersFilter = fmt.Sprintf("manager IN (%s)", managers)
	}

	developersFilter := ""
	if developers != "" {
		developersFilter = fmt.Sprintf("developer IN (%s)", developers)
	}

	prizeMinFilter := ""
	if prizeMin != 0 {
		prizeMinFilter = fmt.Sprintf("prize > %d", prizeMin)
	}

	prizeMaxFilter := ""
	if prizeMax != 0 {
		prizeMaxFilter = fmt.Sprintf("prize < %d", prizeMax)
	}

	startTimeFilter := ""
	if startTime != 0 {
		startTimeFilter = fmt.Sprintf("startTime > %d", startTime)
	}

	endTimeFilter := ""
	if endTime != 0 {
		endTimeFilter = fmt.Sprintf("endTime < %d", endTime)
	}

	queryFilter := ""
	if query != "" {
		queryFilter = fmt.Sprintf("(nameRu LIKE '%%%s%%' OR nameEn LIKE '%%%s%%' OR nameKz LIKE '%%%s%%')", query, query, query)
	}

	allFilters := []string{disciplinesFilter, managersFilter, developersFilter, prizeMinFilter, prizeMaxFilter, startTimeFilter, endTimeFilter, queryFilter}
	var nonEmptyFilters []string
	for _, filter := range allFilters {
		if filter != "" {
			nonEmptyFilters = append(nonEmptyFilters, filter)
		}
	}

	sqlParams := strings.Join(nonEmptyFilters, " AND ")

	where := ""
	if sqlParams != "" {
		where = "WHERE"
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM Events %s %s;", where, sqlParams)

	rows, _ := db.Query(sqlQuery)
	defer rows.Close()

	events := []entities.Event{}

	for rows.Next() {
		e := entities.Event{}
		rows.Scan(&e.Id, &e.PreviewPhoto, &e.Name.Ru, &e.Name.En, &e.Name.Kz, &e.Description.Ru, &e.Description.En, &e.Description.Kz, &e.Manager, &e.Developer, &e.Place.Ru, &e.Place.En, &e.Place.Kz, &e.Discipline, &e.StartTime, &e.EndTime, &e.Prize)
		events = append(events, e)
	}

	return events
}

func GetEventById(id int) entities.Event {
	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Events WHERE id = $1", id)

	event := entities.Event{}
	row.Scan(&event.Id, &event.PreviewPhoto, &event.Name.Ru, &event.Name.En, &event.Name.Kz, &event.Description.Ru, &event.Description.En, &event.Description.Kz, &event.Manager, &event.Developer, &event.Place.Ru, &event.Place.En, &event.Place.Kz, &event.Discipline, &event.StartTime, &event.EndTime, &event.Prize)

	return event
}

func RemoveEvent(id int) {
	db, _ := sql.Open("sqlite3", "store.db")

	defer db.Close()

	db.Exec("DELETE FROM Events WHERE id = $1", id)
}

func GetFilters() []byte {

	db, _ := sql.Open("sqlite3", "store.db")
	defer db.Close()

	rowsManagers, _ := db.Query("SELECT DISTINCT manager FROM Events")
	rowsDisciplines, _ := db.Query("SELECT DISTINCT discipline FROM Events")
	rowsDevelopers, _ := db.Query("SELECT DISTINCT developer FROM Events")
	defer rowsManagers.Close()
	defer rowsDisciplines.Close()
	defer rowsDevelopers.Close()

	managers := []string{}
	disciplines := []string{}
	developers := []string{}

	for rowsManagers.Next() {
		var s string
		rowsManagers.Scan(&s)
		managers = append(managers, s)
	}

	for rowsDisciplines.Next() {
		var s string
		rowsDisciplines.Scan(&s)
		disciplines = append(disciplines, s)
	}

	for rowsDevelopers.Next() {
		var s string
		rowsDevelopers.Scan(&s)
		developers = append(developers, s)
	}

	data := map[string][]string{
		"managers":    managers,
		"disciplines": disciplines,
		"developers":  developers,
	}

	jsonData, _ := json.Marshal(data)

	return jsonData
}
