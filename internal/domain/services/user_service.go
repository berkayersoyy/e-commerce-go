package services

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

//UserService User service
type UserService interface {
	Update(ctx context.Context, user models.User) error
	FindByUUID(ctx context.Context, id string) (models.User, error)
	Insert(ctx context.Context, user models.User) error
	Delete(ctx context.Context, uuid string) error
	FindByUsername(ctx context.Context, username string) ([]models.User, error)
	CreateTable(ctx context.Context) error
}
