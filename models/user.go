package models

import (
	"github.com/matthewhartstonge/argon2"
)

type User struct {
	Id       int    `json:"id_user"`         // Sesuai DB: id_user
	RolesId  int    `json:"roles_id"`        // Sesuai DB: roles_id
	Fullname string `json:"fullname"`        // Sesuai DB: fullname
	Email    string `json:"email"`           // Sesuai DB: email
	Phone    string `json:"phone"`           // Sesuai DB: phone
	Password string `json:"password"`        // Sesuai DB: password
	Address  string `json:"address"`         // Sesuai DB: address
	Picture  string `json:"profile_picture"` // Sesuai DB: profile_picture
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var Users []User
var NextID = 1
var Argon = argon2.DefaultConfig()

func HashPassword(password string) string {
	hash, _ := Argon.HashEncoded([]byte(password))
	return string(hash)
}

func VerifyPassword(hash string, password string) bool {
	match, _ := argon2.VerifyEncoded([]byte(password), []byte(hash))
	return match
}

func CreateUserLogic(input User) (Response, int) {
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
	for i := 0; i < len(Users); i++ {
		if Users[i].Email == input.Email {
			return Response{Status: false, Message: "Email already registered"}, 400
		}
	}

	input.Password = HashPassword(input.Password)
	input.Id = NextID
	NextID++

	// Default Role jika tidak diisi (misal: 2 untuk 'User')
	if input.RolesId == 0 {
		input.RolesId = 2
	}

	Users = append(Users, input)

	return Response{
		Status:  true,
		Message: "User registered successfully",
		Data:    input,
	}, 201
}