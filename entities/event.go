package entities

type Event struct {
	Id           int      `json:"id"`
	PreviewPhoto string   `json:"previewPhoto"`
	Name         I18nText `json:"name"`
	Description  I18nText `json:"description"`
	Manager      string   `json:"manager"`
	Developer    string   `json:"developer"`
	Place        I18nText `json:"place"`
	Discipline   string   `json:"discipline"`
	StartTime    int      `json:"startTime"`
	EndTime      int      `json:"endTime"`
	Prize        int      `json:"prize"`
}
