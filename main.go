package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Picture  string `json:"picture"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var users []User
var nextID = 1

var argon = argon2.DefaultConfig()

func hashPassword(password string) string {
	hash, _ := argon.HashEncoded([]byte(password))
	return string(hash)
}

func verifyPassword(hash string, password string) bool {
	match, _ := argon2.VerifyEncoded([]byte(password), []byte(hash))
	return match
}

func createUserLogic(input User) (Response, int) {
	hasAt := false

	// Basic Validation
	if input.Email == "" || input.Password == "" || input.Fullname == "" {
		return Response{
			Status:  false,
			Message: "Fullname, Email, and Password are required!",
		}, 400
	}

	// Email Validation
	for i := 0; i < len(input.Email); i++ {
		if string(input.Email[i]) == "@" {
			hasAt = true
		}
	}

	if !hasAt {
		return Response{Status: false, Message: "Invalid email format!"}, 400
	}

	if len(input.Password) < 5 {
		return Response{Status: false, Message: "Password must be at least 5 characters!"}, 400
	}

	// Check Duplicate Email
	for i := 0; i < len(users); i++ {
		if users[i].Email == input.Email {
			return Response{Status: false, Message: "Email already registered"}, 400
		}
	}

	input.Password = hashPassword(input.Password)
	input.Id = nextID
	nextID++
	users = append(users, input)

	return Response{
		Status:  true,
		Message: "User registered successfully",
		Data:    input,
	}, 201
}

func main() {
	r := gin.Default()

	// Register
	r.POST("/register", func(ctx *gin.Context) {
		var input User
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(400, Response{
				Status: false, 
				Message: "Invalid request body",
			})
			return
		}
		response, code := createUserLogic(input)
		ctx.JSON(code, response)
	})

	// Login
	r.POST("/login", func(ctx *gin.Context) {
		var input User
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(400, Response{
				Status: false, 
				Message: "Invalid request body",
			})
			return
		}

		for i := 0; i < len(users); i++ {
			if users[i].Email == input.Email {
				if verifyPassword(users[i].Password, input.Password) {
					ctx.JSON(200, Response{
						Status: true, 
						Message: "Login successful", 
						Data: users[i],
					})
					return
				}
				ctx.JSON(401, Response{
					Status: false, 
					Message: "Wrong Password",
				})
				return
			}
		}
		ctx.JSON(404, Response{
			Status: false, 
			Message: "Email not Registered",
		})
	})

	// Get All Users
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Status: true, 
			Message: "Success get all users", 
			Data: users,
		})
	})

	// Get User By ID
	r.GET("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].Id) {
				ctx.JSON(200, Response{
					Status: true, 
					Message: "User Found!", 
					Data: users[i],
				})
				return
			}
		}
		ctx.JSON(404, Response{Status: false, Message: "User not found!"})
	})

	// Update User 
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var input User

		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(400, Response{
				Status: false, 
				Message: "Invalid request body",
			})
			return
		}

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].Id) {
				if input.Fullname != "" {
					users[i].Fullname = input.Fullname
				}
				
				if input.Email != "" && input.Email != users[i].Email {
					for j := 0; j < len(users); j++ {
						if users[j].Email == input.Email {
							ctx.JSON(400, Response{
								Status: false, 
								Message: "Email already Registered"})
							return
						}
					}
					users[i].Email = input.Email
				}

				if input.Phone != "" {
					users[i].Phone = input.Phone
				}

				if input.Password != "" {
					users[i].Password = hashPassword(input.Password)
				}

				if input.Address != "" {
					users[i].Address = input.Address
				}

				if input.Picture != "" {
					users[i].Picture = input.Picture
				}

				ctx.JSON(200, Response{
					Status: true, 
					Message: "User updated successfully!", 
					Data: users[i],
				})
				return
			}
		}
		ctx.JSON(404, Response{Status: false, Message: "User not found"})
	})

	// Delete User
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].Id) {
				users = append(users[:i], users[i+1:]...)
				ctx.JSON(200, Response{
					Status: true, 
					Message: "User Deleted successfully!",
				})
				return
			}
		}
		ctx.JSON(404, Response{Status: false, Message: "User not found"})
	})

	r.Run("localhost:8888")
}
