package models

import "time"

const (
	UserNotFound = "user not found"
)

type User struct {
	UUID      string     `json:"UUID"`
	Username  string     `json:"Username" validate:"required"`
	Password  string     `json:"Password" validate:"required"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}
