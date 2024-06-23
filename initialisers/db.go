package initialisers

import (
	"fmt"
	"os"
	"pagination/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	DB.AutoMigrate(&models.Person{})
}

func ConnectToDB() {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDatabase := os.Getenv("POSTGRES_DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London", postgresHost, postgresDatabase, postgresUser, postgresPort)

	connectedDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil || connectedDB == nil {
		panic("failed to connect database")
	}
	DB = connectedDB
}
