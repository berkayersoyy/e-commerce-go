package user

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/services"
)

//userService user Service
type userService struct {
	UserRepository repositories.UserRepository
}

func (u userService) Update(ctx context.Context, user models.User) error {
	err := u.UserRepository.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u userService) FindByUUID(ctx context.Context, uuid string) (models.User, error) {
	user, err := u.UserRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (u userService) FindByUsername(ctx context.Context, username string) (models.User, error) {
	user, err := u.UserRepository.FindByUsername(ctx, username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u userService) Insert(ctx context.Context, user models.User) error {
	err := u.UserRepository.Insert(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u userService) Delete(ctx context.Context, id string) error {
	err := u.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (u userService) CreateTable(ctx context.Context) error {
	err := u.UserRepository.CreateTable(ctx)
	if err != nil {
		return err
	}
	return nil
}

//ProvideUserService Provide user service
func ProvideUserService(u repositories.UserRepository) services.UserService {
	return userService{UserRepository: u}
}
