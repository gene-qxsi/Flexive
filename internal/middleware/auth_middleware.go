package middleware

import (
	"net/http"
	"strings"

	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	AuthSrv *services.AuthService
}

func NewAuthMiddleware(AuthSrv *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{AuthSrv: AuthSrv}
}

func (m *AuthMiddleware) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "пользователь не прошел аутентификацию"})
			return
		}

		data := strings.Split(token, " ")
		if len(data) != 2 || data[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "неверный формат заголовка Authorization"})
			return
		}

		claims, err := m.AuthSrv.ParseAccessToken(data[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
