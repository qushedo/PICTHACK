package entities

type Course struct {
	Id            uint   `json:"id"`
	University    string `json:"university"`
	Megafaculty   string `json:"megafaculty"`
	MegafacultyId string `json:"megafaculty_id"`
	FacultyId     string `json:"faculty_id"`
	Faculty       string `json:"faculty"`
	Subject       string `json:"subject"`
	Link          string `json:"link"`
	CourseName    string `json:"course_name"`
}
