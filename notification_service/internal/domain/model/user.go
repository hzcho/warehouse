package model

import "github.com/google/uuid"

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Login       string    `json:"login"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}
