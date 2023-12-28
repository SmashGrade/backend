package entity

type State struct {
	Basemodel
	Description string
	ModuleID    uint // used for has one association
}
