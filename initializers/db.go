package initializers

import (
	"fmt"
	"os"

	"github.com/Thalesteo/shortner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Fail to connect database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
