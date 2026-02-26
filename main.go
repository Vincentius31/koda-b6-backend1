package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User
var nextID = 1

func main() {
	r := gin.Default()

	// Create Users
	r.POST("/users", func(c *gin.Context) {
		var input User

		if c.BindJSON(&input) != nil {
			c.JSON(400, gin.H{
				"status":  false,
				"message": "Invalid request body",
				"data":    nil,
			})
			return
		}

		if input.Email == "" || input.Password == "" {
			c.JSON(400, gin.H{
				"status":  false,
				"message": "Email & Password required",
				"data":    nil,
			})
			return
		}

		for i := 0; i < len(users); i++ {
			if users[i].Email == input.Email {
				c.JSON(400, gin.H{
					"status":  false,
					"message": "Email already registered",
					"data":    nil,
				})
				return
			}
		}

		input.ID = nextID
		nextID++

		users = append(users, input)

		c.JSON(201, gin.H{
			"status":  true,
			"message": "User created successfully",
			"data":    input,
		})
	})

	// Get all users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Success get users",
			"data":    users,
		})
	})

	// Get Users by ID
	r.GET("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].ID) {
				c.JSON(200, gin.H{
					"status":  true,
					"message": "User found",
					"data":    users[i],
				})
				return
			}
		}

		c.JSON(404, gin.H{
			"status":  false,
			"message": "User not found",
			"data":    nil,
		})
	})

	// Update Users
	r.PATCH("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		var input User

		if c.BindJSON(&input) != nil {
			c.JSON(400, gin.H{
				"status":  false,
				"message": "Invalid request body",
				"data":    nil,
			})
			return
		}

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].ID) {
				if input.Email != "" {
					for j := 0; j < len(users); j++ {
						if users[j].Email == input.Email && users[j].ID != users[i].ID {
							c.JSON(400, gin.H{
								"status":  false,
								"message": "Email already registered",
								"data":    nil,
							})
							return
						}
					}
					users[i].Email = input.Email
				}

				if input.Password != "" {
					users[i].Password = input.Password
				}

				c.JSON(200, gin.H{
					"status":  true,
					"message": "User updated successfully",
					"data":    users[i],
				})
				return
			}
		}

		c.JSON(404, gin.H{
			"status":  false,
			"message": "User not found",
			"data":    nil,
		})
	})

	// Delete Users
	r.DELETE("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].ID) {

				users = append(users[:i], users[i+1:]...)

				c.JSON(200, gin.H{
					"status":  true,
					"message": "User deleted successfully",
					"data":    nil,
				})
				return
			}
		}

		c.JSON(404, gin.H{
			"status":  false,
			"message": "User not found",
			"data":    nil,
		})
	})

	r.Run("localhost:8888")
}
