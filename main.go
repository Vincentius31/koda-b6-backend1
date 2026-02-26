package main

import "github.com/gin-gonic/gin"

type User struct {
	ID       int    `json:id`
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
		}

		if input.Email == "" || input.Password == "" {
			ctx.JSON(400, gin.H{"error" : "Email & Password required"})
		}

		input.ID = nextId
		nextId++

		users = append(users, input)

		ctx.JSON(201, gin.H{
			"message": "User created!",
			"data": input,
		})
	})

	

	r.Run("localhost:8888")
}
