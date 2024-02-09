package api

import "github.com/SmashGrade/backend/app/db"

// Controller is the base controller for all API controllers
type BaseController struct {
	Provider db.Provider
}

// Constructor for BaseController
func NewBaseController(provider db.Provider) *BaseController {
	return &BaseController{
		Provider: provider,
	}
}
