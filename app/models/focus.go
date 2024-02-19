package models

type Focus struct {
	Basemodel
	FieldID     uint   `json:"-"`
	Field       Field  `json:"field"`
	Description string `json:"description"`
}
