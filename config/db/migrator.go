package db

import "gorm.io/gorm"

func Automigrate(db *gorm.DB) error {
	err := db.Migrator().DropTable()
	if err != nil {
		return err
	}
	err = db.AutoMigrate()
	if err != nil {
		return err
	}

	return nil
}
