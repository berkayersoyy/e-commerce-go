package dto

type ProductDto struct {
	ID          uint    `gorm:"primary_key"`
	Name        string  `json:"name" validate:"required,min=5,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
