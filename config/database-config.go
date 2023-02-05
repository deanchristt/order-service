package config

import (
	"fmt"
	"github.com/deanchristt/order-service/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(
		&entity.Customer{},
		&entity.Seller{},
		&entity.Product{},
		&entity.Transaction{},
		&entity.Payment{},
		&entity.Order{},
		&entity.OrderProduct{},
	)
	return db

}

func CloseDatabaseConnection(db *gorm.DB) {
	dbPostgre, err := db.DB()
	if err != nil {
		panic("Failed to close connection database")
	}
	dbPostgre.Close()
}
