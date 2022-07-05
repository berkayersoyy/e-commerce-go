package dto

//UpdateUserDto User Dto
type UpdateUserDto struct {
	UUID     string `json:"UUID"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type CreateUserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
