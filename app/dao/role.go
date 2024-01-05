package dao

import "github.com/SmashGrade/backend/app/entity"

func (db *Database) ListRoles(roles *[]*entity.Role, searchstring string) error {
	if searchstring != "" {
		db.Db.First(&roles, "description = ?", searchstring)
		return nil
	}

	db.Db.Find(&roles)
	return nil
}
