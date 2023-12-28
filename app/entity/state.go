package entity

type State struct {
	basemodel
	Description string
	ModuleID    uint // used for has one association
}
