package entity

type StudyStage struct {
	Basemodel
	Description string
	ModuleID    uint // used for has one association
}
