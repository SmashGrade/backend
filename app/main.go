package main

import (
	"github.com/SmashGrade/backend/app/cmd"
	_ "github.com/SmashGrade/backend/app/docs"
	_ "gorm.io/gorm"
)

// @title						Smashgrade Backend API
// @version					1.0
// @description				Backend API for Smashgrade, a web application for tracking your student grades.
// @termsOfService				http://swagger.io/terms/
// @contact.name				HFTM Grenchen
// @contact.url				https://www.hftm.ch
// @license.name				Affero General Public License
// @license.url				https://www.gnu.org/licenses/agpl-3.0.html
// @host						api.smashgrade.ch
// @BasePath					/v1
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @accept						json
// @produce					json
// @schemes					https
// @description				Type "Bearer" followed by a space and JWT token.
// @contact.name				Backend Support
// @contact.email				backend@smashgrade.ch
func main() {
	cmd.Execute()
}
