package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"QLLHTT/internal/models" // 🔁 Cập nhật đúng tên module bạn dùng
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Load .env failed")
	}
}

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect DB:", err)
	}

	// ✅ Auto migrate các bảng
	err = DB.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Document{},
		&models.Enrollment{},
	)
	if err != nil {
		log.Fatal("❌ Auto migration failed:", err)
	}

	fmt.Println("✅ Connected to DB and Migrated!")
}
