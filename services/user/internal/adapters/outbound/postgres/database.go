package postgres

import (
	"fmt"

	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/config/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	conf := database.NewPostgresConfig()

	// Initialize the database connection
	return gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%t TimeZone=Asia/Jakarta",
			conf.Host, conf.User, conf.Password, conf.DBName, conf.Port, conf.SSLMode),
	}))
}
