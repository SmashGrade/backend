package dao

import "github.com/SmashGrade/backend/legacy/entity"

// returns a field
func (db *Database) GetFieldByDescription(name string) (field *entity.Field, err error) {
	field = &entity.Field{}
	err = db.Db.Model(&entity.Field{}).Where("description = ?", name).First(&field).Error // do we need Model(&entity.Field{}).?
	return
}

// returns a field by id
func (db *Database) GetFieldById(id uint) (field *entity.Field, err error) {
	field = &entity.Field{}
	err = db.Db.First(&field, id).Error
	return
}

// used for testing create field
func (db *Database) CreateField(description string) (field *entity.Field, err error) {
	field = &entity.Field{}
	field.Description = description
	err = db.Db.Create(&field).Error
	return
}
