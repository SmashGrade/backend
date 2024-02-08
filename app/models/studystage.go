package models

type StudyStage struct {
	Basemodel
	Description string `json:"description"`
	ModuleID    uint   `json:"moduleID"`
}
