package dao

import "github.com/SmashGrade/backend/app/entity"

// returns a focus by name
func (db *Database) GetFocusByDescription(name string) (focus *entity.Focus, err error) {
	focus = &entity.Focus{}
	err = db.Db.Model(&entity.Focus{}).Where("description = ?", name).First(&focus).Error
	return
}

// used for testing create focus
func (db *Database) CreateFocus(description string, fieldId uint) (focus *entity.Focus, err error) {
	focus = &entity.Focus{}
	focus.FieldID = fieldId
	focus.Description = description
	err = db.Db.Create(&focus).Error
	return
}
