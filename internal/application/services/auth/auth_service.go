package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/sethvargo/go-retry"
	"github.com/twinj/uuid"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//authService Auth service
type authService struct {
	Client *redis.Client
}

// ExtractToken Extract token
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (a *authService) ValidateToken(r *http.Request) (*jwt.Token, error) {
	accessSecret := os.Getenv("ACCESS_SECRET")
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (a *authService) TokenValid(r *http.Request) error {
	token, err := a.ValidateToken(r)
	if err != nil {
		return err
	}
	claims := make(jwt.MapClaims)
	if err := claims.Valid(); err != nil || !token.Valid {
		return err
	}
	return nil
}
func (a *authService) CreateToken(userid string) (*models.TokenDetails, error) {
	accessSecret := os.Getenv("ACCESS_SECRET")
	refreshSecret := os.Getenv("REFRESH_SECRET")
	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = td.AccessUUID + "++" + userid

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	td.AccessToken, err = at.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (a *authService) CreateAuth(userid string, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := a.Client.Set(td.AccessUUID, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := a.Client.Set(td.RefreshUUID, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
func (a *authService) ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	token, err := a.ValidateToken(r)
	if err != nil {
		return nil, err
	}
	claims := make(jwt.MapClaims)
	if err := claims.Valid(); err != nil && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID := fmt.Sprintf("%.f", claims["user_id"])
		if err != nil {
			return nil, err
		}
		return &models.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}
func (a *authService) DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := a.Client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
func (a *authService) DeleteTokens(authD *models.AccessDetails) error {
	refreshUUID := fmt.Sprintf("%s++%s", authD.AccessUUID, authD.UserID)
	deletedAt, err := a.Client.Del(authD.AccessUUID).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := a.Client.Del(refreshUUID).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("internal error")
	}
	return nil
}
func (a *authService) FetchAuth(authD *models.AccessDetails) (uint64, error) {
	userid, err := a.Client.Get(authD.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	if authD.UserID != userid {
		return 0, errors.New("unauthorized")
	}
	return userID, nil
}

//initRedis Init redis
func initRedis() *redis.Client {
	dsn := os.Getenv("REDIS_HOST")
	var client *redis.Client
	ctx := context.Background()
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		client = redis.NewClient(&redis.Options{
			Addr: dsn,
		})
		if _, err := client.Ping().Result(); err != nil {
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return client
}

// ProvideAuthService Provides auth service
func ProvideAuthService() services.AuthService {
	return &authService{Client: initRedis()}
}
