package routes

import (
	"koda-b6-backend1/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Routes
	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/users", handlers.GetAllUsersHandler)
	r.GET("/users/:id", handlers.GetUserByIdHandler)
	r.PATCH("/users/:id", handlers.UpdateUserHandler)
	r.DELETE("/users/:id", handlers.DeleteUserHandler)

	//Product Routes
	r.POST("/products", handlers.CreateProductHandler)
	r.GET("/products", handlers.GetAllProductsHandler)
	r.GET("/products/:id", handlers.GetProductByIdHandler)
	r.PATCH("/products/:id", handlers.UpdateProductHandler)
	r.DELETE("/products/:id", handlers.DeleteProductHandler)
}