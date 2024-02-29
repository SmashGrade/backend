package models

type MetaCourse struct {
	Teachers  []User     `json:"teachers"`
	Modules   []Module   `json:"modules"`
	Examtypes []Examtype `json:"examtypes"`
}
