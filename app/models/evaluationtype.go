package models

type Evaluationtype struct {
	Basemodel
	Description string `json:"description"`
	Code        string `json:"code"`
}
