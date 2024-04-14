package entities

type User struct {
	Id           int64  `json:"id" gorm:"unique"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"user_name"`
	Localisation string `json:"localisation"`
	University   string `json:"university"`
	Faculty      string `json:"faculty"`
	Items        string `json:"items"`
}
