package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var G_db *gorm.DB

func ConnDB() {
	dbName := os.Getenv("DB_NAME")
	dbLocal := os.Getenv("DB_LOCAL")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbLocal, dbName)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect database failed")
	}
	client.AutoMigrate()
	G_db = client
}
