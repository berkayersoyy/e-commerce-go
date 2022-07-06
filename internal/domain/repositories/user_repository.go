package repositories

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

//UserRepository User repository
type UserRepository interface {
	Update(ctx context.Context, user models.User) error
	FindByUUID(ctx context.Context, id string) (models.User, error)
	Insert(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id string) error
	FindByUsername(ctx context.Context, username string) ([]models.User, error)
	CreateTable(ctx context.Context) error
}
