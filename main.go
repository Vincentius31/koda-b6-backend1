package main

import (
	"os"
	"koda-b6-backend1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	r := gin.Default()
	r.Static("/uploads", "./uploads")
	
	routes.SetupRoutes(r)

	r.Run("localhost:8888")
}