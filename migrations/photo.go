package migrations

import (
	"log"

	"github.com/farrelahmady/my-gram-hacktiv8/models"
)

func (m Migration) PhotoMigrate() {
	model := &models.Photo{}
	if m.db.Migrator().HasTable(&models.Photo{}) {

		err := m.db.Migrator().DropTable(&model)
		if err != nil {
			log.Fatal("Failed to drop ", model.TableName(), "table, error:", err)
		}
	}
	err := m.db.AutoMigrate(&models.Photo{})
	if err != nil {
		log.Fatal("Failed to migrate ", model.TableName(), "table, error:", err)
	}
}
