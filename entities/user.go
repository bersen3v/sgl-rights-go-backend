package entities

type User struct {
	Id           int    `json:"id"`
	PreviewPhoto string `json:"previewPhoto"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Company      string `json:"company"`
	Mail         string `json:"mail"`
	Phone        string `json:"phone"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	IsAdmin      int    `json:"isAdmin"`
}
