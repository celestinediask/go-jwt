package initializers

import "github.com/celestinediask/go-jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
