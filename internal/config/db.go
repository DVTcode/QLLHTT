package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" //Dùng để tải file .env vào biến môi trường trong Go.
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"QLLHTT/internal/models" // 🔁 Cập nhật đúng tên module bạn dùng
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load() //Load để đọc file .env, không truyền vào gì thì nó tự động mặc định kiếm .env để gán vào biến môi trường và nếu file .env bị lỗi thì trả về false nghĩa là ko phải nil
	if err != nil {
		log.Fatal("❌ Load .env failed") //Dùng để ghi log dạng lỗi nghiêm trọng rồi thoát khỏi chương trình ngay lập tức (giống như panic()).
	}
}

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", // sau khi format root:123456@tcp(localhost:3306)/qllhtt?charset=utf8mb4&parseTime=True
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //mở kết nối đến DB qua driver mysql.
	if err != nil {
		log.Fatal("❌ Failed to connect DB:", err)
	}

	// ✅ Auto migrate các bảng
	err = DB.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Document{},
		&models.Enrollment{},
		&models.RefreshToken{},
	)
	if err != nil {
		log.Fatal("❌ Auto migration failed:", err)
	}

	fmt.Println("✅ Connected to DB and Migrated!")
}
