package schemas

type User struct {
	Id             int64  `json:"id"`
	ClassStartYear int64  `json:"classStartYear"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}
