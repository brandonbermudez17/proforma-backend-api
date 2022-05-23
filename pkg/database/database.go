package database

import (
	"proforma-backend-api/pkg/configuration"
	"proforma-backend-api/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dbString := configuration.GetConnectionString()
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	migrate(db)
	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.PersonalData{},
		&models.User{},
		&models.Company{},
		&models.Client{},
		&models.Representative{},
		&models.Supplier{},
	)
}
