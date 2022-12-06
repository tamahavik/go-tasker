package database

import (
	"log"
	"project-tasker/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	connStr := `postgres://postgres:password@localhost:5432/tasker?sslmode=disable`
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln("FAILED TO CONNECT DATABASE")
		panic(err)
	}

	db.AutoMigrate(&models.Customer{})

	return db
}
