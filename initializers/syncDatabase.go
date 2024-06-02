package initializers

import "github.com/SoroushBeigi/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}