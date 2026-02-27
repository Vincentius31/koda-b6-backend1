package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
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


func createUserLogic(input User) (Response, int){
	hasAt := false // Untuk validasi email

	if input.Email == "" || input.Password == "" {
		return Response{
			Status: false,
			Message: "Email and Password required!",
			Data: nil,
		},400
	}

	// Validasi Email

	for i := 0; i < len(input.Email); i++ {
		if string(input.Email[i]) == "@" {
			hasAt = true
		}
	}

	if !hasAt {
		return Response{
			Status: false,
			Message: "Invalid email format!",
			Data: nil,
		},400
	}

	// Validasi Password minimal 5 karakter

	if len(input.Password) < 5 {
		return Response{
			Status: false,
			Message: "Password must be at least 5 characters!",
			Data: nil,
		},400
	}

	// Cek Duplicate Email

	for i := 0; i < len(users); i++ {
		if users[i].Email == input.Email {
			return Response{
				Status: false,
				Message: "Email already registered",
				Data: nil,
			},400
		}
	}

	input.Password = hashPassword(input.Password)

	input.ID = nextID
	nextID++
	users = append(users, input)

	return Response{
		Status: true,
		Message: "User registered successfully",
		Data: input,
	}, 201
}

func main() {
	r := gin.Default()

	// Create Users
	r.POST("/users", func(ctx *gin.Context) {
		var input User

		if ctx.BindJSON(&input) != nil {
			ctx.JSON(400, Response{
				Status: false,
				Message: "Invalid request body!",
				Data: nil,
			})
			return
		}

		response, code := createUserLogic(input)
		ctx.JSON(code, response)
	})

	// Register Users
	r.POST("/register", func(ctx *gin.Context) {
		var input User

		if ctx.BindJSON(&input) != nil {
			ctx.JSON(400, Response{
				Status: false,
				Message: "Invalid request body",
				Data: nil,
			})
			return
		}

		response, code := createUserLogic(input)
		ctx.JSON(code, response)
	})

	//Login Users
	r.POST("/login", func(ctx *gin.Context) {
		var input User

		if ctx.BindJSON(&input) != nil {
			ctx.JSON(400, Response{
				Status: false,
				Message: "Invalid request body",
				Data: nil,
			})
			return
		}

		if input.Email == "" || input.Password == "" {
			ctx.JSON(400, Response{
				Status: false,
				Message: "Email and Password Required!",
				Data: nil,
			})
			return
		}

		for i := 0; i < len(users); i++ {
			if users[i].Email == input.Email {
				if verifyPassword(users[i].Password, input.Password) {
					ctx.JSON(200, Response{
						Status: true,
						Message: "Login successfuly",
						Data: users[i],
					})
					return
				}

				ctx.JSON(401, Response{
					Status: false,
					Message: "Wrong Password",
					Data: nil,
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Status: false,
			Message: "Email not Registered",
			Data: nil,
		})
	})

	// Get all users
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Status: true,
			Message: "Success get all users",
			Data: users,
		})
	})

	// Get Users by ID
	r.GET("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		for i := 0; i < len(users); i++ {
			if idParam == fmt.Sprint(users[i].ID) {
				ctx.JSON(200, Response{
					Status: true,
					Message: "User Found!",
					Data: users[i],
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Status: false,
			Message: "User not found!",
			Data: nil,
		})
	})

	// Update Users
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var input User

		if ctx.BindJSON(&input) != nil {
			ctx.JSON(400, Response{
				Status: false,
				Message: "Invalid request body",
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
								Status: false,
								Message: "Email already Registered",
								Data: nil,
							})
							return
						}
					}
					users[i].Email = input.Email
				}

				if input.Password != "" {
					users[i].Password = hashPassword(input.Password)
				}

				ctx.JSON(200, Response{
					Status: true,
					Message: "User updated succsessfully!",
					Data: users[i],
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Status: false,
			Message: "User not found",
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
					Status: true,
					Message: "User Deleted succsessfully!",
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Status: false,
			Message: "User not found",
			Data: nil,
		})
	})

	r.Run("localhost:8888")
}


