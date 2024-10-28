package response

import "github.com/google/uuid"

type Users struct {
	Users []BaseUserInfo `json:"users"`
}

type BaseUserInfo struct {
	ID          uuid.UUID `json:"id"`
	Login       string    `json:"login"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}
