package handlers

import (
	"fmt"
	"koda-b6-backend1/models"
	"github.com/gin-gonic/gin"
)

// CreateHandler godoc
// @Summary      Create a new user
// @Description  Register a new user into the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserRegisterDTO  true  "Register Data"
// @Success      201   {object}  models.Response
// @Failure      400   {object}  models.Response
// @Router       /users [post]
func CreateHandler(ctx *gin.Context) {
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

// LoginHandler godoc
// @Summary      User Login
// @Description  Authenticate user with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.UserLoginDTO  true  "Login Credentials"
// @Success      200          {object}  models.Response
// @Failure      401          {object}  models.Response
// @Failure      404          {object}  models.Response
// @Router       /login [post]
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
				Message: "Wrong Email or Password",
			})
			return
		}
	}
	ctx.JSON(404, models.Response{
		Status:  false,
		Message: "Wrong Email or Password",
	})
}

// GetAllUsersHandler godoc
// @Summary      Get all users
// @Description  Retrieve a list of all registered users
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.Response
// @Router       /users [get]
func GetAllUsersHandler(ctx *gin.Context) {
	ctx.JSON(200, models.Response{
		Status:  true,
		Message: "Success get all users",
		Data:    models.Users,
	})
}

// GetUserByIdHandler godoc
// @Summary      Get user by ID
// @Description  Retrieve single user details by their ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /users/{id} [get]
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

// UpdateUserHandler godoc
// @Summary      Update user
// @Description  Update existing user information by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      string       true  "User ID"
// @Param        user  body      models.User  true  "Updated User Data"
// @Success      200   {object}  models.Response
// @Failure      400   {object}  models.Response
// @Failure      404   {object}  models.Response
// @Router       /users/{id} [put]
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

// DeleteUserHandler godoc
// @Summary      Delete user
// @Description  Remove a user from the system by ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /users/{id} [delete]
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