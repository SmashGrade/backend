package dao

import "github.com/SmashGrade/backend/legacy/entity"

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

// returns role by description
func (db *Database) GetRoleByName(role *entity.Role, description string) error {
	return db.Db.First(&role, "description = ?", description).Error
}
