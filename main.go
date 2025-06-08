package main

import (
	"net/http"
	eventApi "sgl-rights/api/event"
	photoApi "sgl-rights/api/photo"
	saleApi "sgl-rights/api/sale"
	userApi "sgl-rights/api/user"

	"sgl-rights/db"

	"github.com/rs/cors"
)

func main() {
	db.CreateDb()

	mux := http.NewServeMux()

	mux.HandleFunc("/createEvent", eventApi.CreateEvent)
	mux.HandleFunc("/getAllEvents", eventApi.GetAllEvents)
	mux.HandleFunc("/getEventById", eventApi.GetEventById)
	mux.HandleFunc("/removeEvent", eventApi.RemoveEvent)
	mux.HandleFunc("/updateEvent", eventApi.UpdateEvent)
	mux.HandleFunc("/searchEvents", eventApi.SearchEvents)
	mux.HandleFunc("/getEventsFilters", eventApi.GetEventsFilters)
	mux.HandleFunc("/getPhoto", photoApi.GetPhoto)

	mux.HandleFunc("/createUser", userApi.CreateUser)
	mux.HandleFunc("/getAllUsers", userApi.GetAllUsers)
	mux.HandleFunc("/getUserById", userApi.GetUserById)
	mux.HandleFunc("/updateUser", userApi.UpdateUser)
	mux.HandleFunc("/removeUser", userApi.RemoveUser)
	mux.HandleFunc("/authUser", userApi.AuthUser)
	mux.HandleFunc("/addEventToUser", userApi.AddEventToUser)
	mux.HandleFunc("/removeEventFromUser", userApi.RemoveEventFromUser)
	mux.HandleFunc("/getUserEvents", userApi.GetUserEvents)
	mux.HandleFunc("/getAllSales", saleApi.GetAllSales)

	corsHandler := cors.Default().Handler(mux)

	http.Handle("/", corsHandler)
	http.ListenAndServe(":8090", nil)
}

// func main() {
// 	fmt.Println(db.GetAllEvents())
// }
