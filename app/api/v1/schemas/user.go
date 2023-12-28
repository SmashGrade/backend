package schemas

type User struct {
	Id                  int    `json:"id"`
	CurriculumStartYear int    `json:"curriculumStartYear"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	Role                string `json:"role"`
}
