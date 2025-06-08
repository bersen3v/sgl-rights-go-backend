package entities

type Sale struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
	Time    int `json:"time"`
}
