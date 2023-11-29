package schemas

type User struct {
	Id             int    `json:"id"`
	ClassStartYear int    `json:"classStartYear"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}
