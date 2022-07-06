package dto

import (
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/twinj/uuid"
	"time"
)

//UpdateUserDto Update User Dto
type UpdateUserDto struct {
	UUID     string `json:"UUID"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//CreateUserDto Create User Dto
type CreateUserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateToUser(userDto CreateUserDto) models.User {
	return models.User{
		UUID:      uuid.NewV4().String(),
		Username:  userDto.Username,
		Password:  userDto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
