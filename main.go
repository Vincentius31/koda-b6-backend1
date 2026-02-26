package main

import "github.com/gin-gonic/gin"

type User struct {
	ID      int    `json:id`
	Email    string `json:email`
	Password string `json:password`
}

var users []User
var nextId = 1

func main() {
	r := gin.Default()

	// Create user
	r.POST("/users", func(ctx *gin.Context) {
		var input User

		if ctx.BindJSON(&input) != nil {
			ctx.JSON(400, gin.H{"error" : "Invalid request"})
			return
		}

		if input.Email == "" || input.Password == "" || input.Email == " " || input.Password == " " {
			ctx.JSON(400, gin.H{"error" : "Email & Password required"})
			return
		}

		input.ID = nextId
		nextId++

		users = append(users, input)

		ctx.JSON(201, gin.H{
			"message": "User created!",
			"data": input,
		})
	})

	// Read all user
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Success get users",
			"data": users,
		})
	})

	// Read by Id
	

	r.Run("localhost:8888")
}
