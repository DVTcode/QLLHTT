package main

import (
	"log"
	"os"

	"QLLHTT/internal/config"
	"QLLHTT/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()       //khởi tạo một router mặc định,
	routes.RegisterRoutes(r) //gọi hàm RegisterRoutes trong package routes và truyền router r vào để đăng ký toàn bộ các route (đường dẫn API) của ứng dụng web.

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	log.Println("🚀 Server running at http://localhost" + port)
	r.Run(port)
}
