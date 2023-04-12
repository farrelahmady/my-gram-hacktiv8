package migrations

import (
	"log"
	"reflect"

	"gorm.io/gorm"
)

type Migration struct {
	db *gorm.DB
}

func Execute(db *gorm.DB, migrationMethodName ...string) {
	m := Migration{db}

	migrationType := reflect.TypeOf(m)

	// Execute all seeders if no method name is given
	if len(migrationMethodName) == 0 {
		log.Println("Running all Migration...")
		// We are looping over the method on a Seed struct
		for i := 0; i < migrationType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := migrationType.Method(i)
			// Execute seeder
			migrate(m, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range migrationMethodName {
		migrate(m, item)
	}
}
func migrate(m Migration, migrationMethodName string) {
	// Get the reflect value of the method
	method := reflect.ValueOf(m).MethodByName(migrationMethodName)
	// Exit if the method doesn't exist
	if !method.IsValid() {
		log.Fatal("No method called ", migrationMethodName)
	}
	// Execute the method
	log.Println("Run", migrationMethodName)
	method.Call(nil)

	log.Println(migrationMethodName, "succeed")
}
