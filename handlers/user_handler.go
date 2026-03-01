package handlers

import (
	"fmt"
	"koda-b6-backend1/models"
	"github.com/gin-gonic/gin"
)

// Register
func RegisterHandler(ctx *gin.Context) {
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}
	response, code := models.CreateUserLogic(input)
	ctx.JSON(code, response)
}

// Login
func LoginHandler(ctx *gin.Context) {
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	for i := 0; i < len(models.Users); i++ {
		if models.Users[i].Email == input.Email {
			if models.VerifyPassword(models.Users[i].Password, input.Password) {
				ctx.JSON(200, models.Response{
					Status:  true,
					Message: "Login successful",
					Data:    models.Users[i],
				})
				return
			}
			ctx.JSON(401, models.Response{
				Status:  false,
				Message: "Wrong Password",
			})
			return
		}
	}
	ctx.JSON(404, models.Response{
		Status:  false,
		Message: "Email not Registered",
	})
}

// Get All Users
func GetAllUsersHandler(ctx *gin.Context) {
	ctx.JSON(200, models.Response{
		Status:  true,
		Message: "Success get all users",
		Data:    models.Users,
	})
}

// Get User By ID
func GetUserByIdHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	for i := 0; i < len(models.Users); i++ {
		if idParam == fmt.Sprint(models.Users[i].Id) {
			ctx.JSON(200, models.Response{
				Status:  true,
				Message: "User Found!",
				Data:    models.Users[i],
			})
			return
		}
	}
	ctx.JSON(404, models.Response{Status: false, Message: "User not found!"})
}

// Update User
func UpdateUserHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var input models.User

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	for i := 0; i < len(models.Users); i++ {
		if idParam == fmt.Sprint(models.Users[i].Id) {
			if input.Fullname != "" {
				models.Users[i].Fullname = input.Fullname
			}

			if input.Email != "" && input.Email != models.Users[i].Email {
				for j := 0; j < len(models.Users); j++ {
					if models.Users[j].Email == input.Email {
						ctx.JSON(400, models.Response{
							Status:  false,
							Message: "Email already Registered"})
						return
					}
				}
				models.Users[i].Email = input.Email
			}

			if input.Phone != "" {
				models.Users[i].Phone = input.Phone
			}

			if input.Password != "" {
				models.Users[i].Password = models.HashPassword(input.Password)
			}

			if input.Address != "" {
				models.Users[i].Address = input.Address
			}

			if input.Picture != "" {
				models.Users[i].Picture = input.Picture
			}

			if input.RolesId != 0 {
				models.Users[i].RolesId = input.RolesId
			}

			ctx.JSON(200, models.Response{
				Status:  true,
				Message: "User updated successfully!",
				Data:    models.Users[i],
			})
			return
		}
	}
	ctx.JSON(404, models.Response{Status: false, Message: "User not found"})
}

// Delete User
func DeleteUserHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	for i := 0; i < len(models.Users); i++ {
		if idParam == fmt.Sprint(models.Users[i].Id) {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			ctx.JSON(200, models.Response{
				Status:  true,
				Message: "User Deleted successfully!",
			})
			return
		}
	}
	ctx.JSON(404, models.Response{Status: false, Message: "User not found"})
}