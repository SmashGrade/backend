package dao

import "github.com/SmashGrade/backend/app/entity"

func (db *Database) ListRoles(roles *[]*entity.Role, searchstrings []string) error {

	for _, searchstring := range searchstrings {
		var role entity.Role
		if searchstring != "" {
			db.Db.First(&role, "description = ?", searchstrings)
			*roles = append(*roles, &role)
		}
	}

	db.Db.Find(&roles)
	return nil
}
