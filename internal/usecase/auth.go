package usecase

import (
	"context"
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gene-qxsi/Flexive/internal/dto"
	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/redis/go-redis/v9"
)

type AuthUseCase struct {
	UserSrv    *services.UserService
	ProfileSrv *services.ProfileService
	AuthSrv    *services.AuthService
}

func NewAuthUseCase(UserSrv *services.UserService, AuthSrv *services.AuthService, ProfileSrv *services.ProfileService) *AuthUseCase {
	return &AuthUseCase{
		UserSrv:    UserSrv,
		AuthSrv:    AuthSrv,
		ProfileSrv: ProfileSrv,
	}
}

func (a *AuthUseCase) SignIn(ctx context.Context, req dto.SignInRequest) (*dto.TokenResponse, error) {
	const op = "internal/usecase/auth_usecase.go/SignIn()"

	user, err := a.UserSrv.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	ok := a.UserSrv.Hasher.Compare(req.Password, user.PasswordHash)
	if !ok {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", "не верный пароль", op)
	}

	accessToken, err := a.AuthSrv.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	refreshToken, err := a.AuthSrv.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	err = a.AuthSrv.SaveRefreshToken(ctx, refreshToken, user.ID)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	tokenResponse := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokenResponse, nil
}

func (a *AuthUseCase) SignUp(ctx context.Context, req dto.SignUpRequest) (*dto.TokenResponse, error) {
	const op = "internal/usecase/auth_usecase.go/SignUp()"

	userDTO, err := a.UserSrv.FindByEmail(ctx, req.Email)
	if err == nil && userDTO != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", "пользователь с таким email уже существует", op)
	}

	user, err := a.UserSrv.CreateUser(&domain.User{
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
	})

	if err != nil {
		return nil, fmt.Errorf("не удалось создать юзера. ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	_, err = a.ProfileSrv.CreateProfile(ctx, domain.Profile{
		UserID:   user.ID,
		Username: req.Username,
		Bio:      "стандартное описание",
		Birthday: nil,
		Website:  nil,
		Role:     nil,
		// AvatarURL: "/default-avatar",
	})
	if err != nil {
		return nil, fmt.Errorf("не удалось создать профиль. ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	accessToken, err := a.AuthSrv.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("неудачная генерация access токена. ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	refreshToken, err := a.AuthSrv.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("неудачная генерация refresh токена. ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	err = a.AuthSrv.SaveRefreshToken(ctx, refreshToken, user.ID)
	if err != nil {
		return nil, fmt.Errorf("сохранение refresh токена ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	tokenResponse := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokenResponse, nil
}

func (a *AuthUseCase) RefreshToken(ctx context.Context, req dto.RefreshToken) (*dto.TokenResponse, error) {
	const op = "internal/usecase/auth_usecase.go/RefreshToken()"

	userID, err := a.AuthSrv.GetUserIDByRefreshToken(ctx, req.RefreshToken)
	if err == redis.Nil {
		return nil, fmt.Errorf("refresh token не найден в Redis")
	}

	if err != nil {
		return nil, fmt.Errorf("ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	user, err := a.UserSrv.GetUser(userID)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
	}

	accessToken, err := a.AuthSrv.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
	}

	tokenResponse := dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: req.RefreshToken,
	}

	return &tokenResponse, nil
}

func (a *AuthUseCase) SignOut(ctx context.Context, req dto.RefreshToken) error {
	const op = "internal/usecase/auth_usecase.go/SignUp()"

	err := a.AuthSrv.DeleteRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	return nil
}
