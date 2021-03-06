package usecases

import (
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

//AuthService Auth service
type AuthService interface {
	ValidateToken(r *http.Request) (*jwt.Token, error)
	TokenValid(r *http.Request) error
	CreateToken(userID string) (*models.TokenDetails, error)
	CreateAuth(userID string, td *models.TokenDetails) error
	DeleteTokens(authD *models.AccessDetails) error
	ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error)
	FetchAuth(authD *models.AccessDetails) (uint64, error)
	DeleteAuth(givenUUID string) (int64, error)
}
