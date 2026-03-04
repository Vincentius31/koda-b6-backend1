package main

import (
	"koda-b6-backend1/routes"
	"github.com/gin-gonic/gin"
)


//@title 		koda-b6-backend1
//@version		1.0.0
//@description	This is basic backend CRUD for users and products
//@host			localhost:8888
//@BasePath		/

func main() {
	r := gin.Default()
	
	routes.SetupRoutes(r)

	r.Run("localhost:8888")
}