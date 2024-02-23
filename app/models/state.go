package models

type State struct {
	Basemodel
	Description string `json:"description"`
	ModuleID    uint   `json:"moduleID"`
}
