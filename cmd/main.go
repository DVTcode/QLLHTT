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

	r := gin.Default()
	routes.RegisterRoutes(r)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	log.Println("🚀 Server running at http://localhost" + port)
	r.Run(port)
}
