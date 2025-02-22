package services

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey  []byte
	refreshTTL time.Duration
)

type AuthClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type AuthService struct {
	Repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func Init() error {
	const op = "internal/services/aut_service.go/InitSecretKey()"

	ttl, err := strconv.Atoi(os.Getenv("REFRESH_TTL"))
	if err != nil {
		return err
	}
	refreshTTL = time.Duration(ttl)

	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		return fmt.Errorf("❌ InitJWT-ОШИБКА-1: %s. ПУТЬ: %s", "JWT_SECRET_KEY не найден в переменных окружения", op)
	}
	return nil
}

func (a *AuthService) GenerateAccessToken(userID int, username, role string) (string, error) {
	const op = "internal/services/aut_service.go/GenerateToken()"
	claims := &AuthClaims{
		ID:       userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(userID),
			Audience:  jwt.ClaimStrings{"frontend", "backend"},
			ID:        strconv.Itoa(userID),
			Issuer:    "flexive.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signingToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("❌ JWT-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return signingToken, err
}

func (a *AuthService) GenerateRefreshToken(userID int, username string) (string, error) {
	const op = "internal/services/aut_service.go/GenerateToken()"

	claims := &AuthClaims{
		ID:       userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signingToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("❌ JWT-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return signingToken, err
}

func (a *AuthService) GenerateTokens(userID int, username, role string) (*dto.TokenResponse, error) {
	const op = "internal/services/aut_service.go/GenerateTokens()"

	accessToken, err := a.GenerateAccessToken(userID, username, role)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	refreshToken, err := a.GenerateRefreshToken(userID, username)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	tokenResponse := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokenResponse, nil
}

func (a *AuthService) ParseToken(tokenString string) (*AuthClaims, error) {
	const op = "internal/services/aut_service.go/ParseToken()"

	var claims AuthClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("❌ JWT-ОШИБКА-1: %s. ПУТЬ: %s", "не верный метод подписи", op)
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("❌ JWT-ОШИБКА-1: %s. ПУТЬ: %s", "недействительный токен", op)
	}

	return &claims, nil
}

func (a *AuthService) SaveRefreshToken(ctx context.Context, token string, userID int) error {
	return a.Repo.SaveRefreshToken(ctx, token, userID)
}

func (a *AuthService) GetUserIDByRefreshToken(ctx context.Context, token string) (int, error) {
	return a.Repo.GetUserIDByRefreshToken(ctx, token)
}

func (a *AuthService) DeleteRefreshToken(ctx context.Context, token string) error {
	return a.Repo.DeleteByRefreshToken(ctx, token)
}
