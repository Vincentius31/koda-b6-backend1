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

type Response struct {
	Message string `json:"message"`
	Status bool `json:"status"`
	Data interface{} `json:"data"`
}

var users []User
var nextID = 1

func main() {
	r := gin.Default()

	// Create Users
	r.POST("/users", func(ctx *gin.Context) {
		var input User

		if ctx.BindJSON(&input) != nil {
			ctx.JSON(400, Response{
				Message: "Invalid request body!",
				Status: false,
				Data: nil,
			})
			return
		}

		if input.Email == "" || input.Password == "" {
			ctx.JSON(400, Response{
				Message: "Email and Password Required!",
				Status: false,
				Data: nil,
			})
			return
		}

		for i := 0; i < len(users); i++ {
			if users[i].Email == input.Email {
				ctx.JSON(400, Response{
					Message: "Email already Registered",
					Status: false,
					Data: nil,
				})
				return
			}
		}

		input.ID = nextID
		nextID++

		users = append(users, input)

		ctx.JSON(201, Response{
			Message: "User Created Succsessfully",
			Status: true,
			Data: input,
		})
	})

	// Get all users
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Message: "Success get all users",
			Status: true,
			Data: users,
		})
	})

	// Get Users by ID
	r.GET("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].ID) {
				ctx.JSON(200, Response{
					Message: "User Found!",
					Status: true,
					Data: users[i],
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Message: "User not found!",
			Status: false,
			Data: nil,
		})
	})

	// Update Users
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var input User

		if ctx.BindJSON(&input) != nil {
			ctx.JSON(400, Response{
				Message: "Invalid request body",
				Status: false,
				Data: nil,
			})
			return
		}

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].ID) {
				if input.Email != "" {
					for j := 0; j < len(users); j++ {
						if users[j].Email == input.Email && users[j].ID != users[i].ID {
							ctx.JSON(400, Response{
								Message: "Email already Registered",
								Status: false,
								Data: nil,
							})
							return
						}
					}
					users[i].Email = input.Email
				}

				if input.Password != "" {
					users[i].Password = input.Password
				}

				ctx.JSON(200, Response{
					Message: "User updated succsessfully!",
					Status: true,
					Data: users[i],
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Message: "User not found",
			Status: false,
			Data: nil,
		})
	})

	// Delete Users
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].ID) {

				users = append(users[:i], users[i+1:]...)

				ctx.JSON(200, Response{
					Message: "User Deleted succsessfully!",
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Message: "User not found",
			Status: false,
			Data: nil,
		})
	})

	r.Run("localhost:8888")
}
