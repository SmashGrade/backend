package schemas

type User struct {
	Id                  uint     `json:"id"`
	CurriculumStartYear int      `json:"curriculumStartYear"`
	Name                string   `json:"name"`
	Email               string   `json:"email"`
	Roles               []string `json:"roles"`
}
