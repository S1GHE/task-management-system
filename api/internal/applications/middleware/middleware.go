package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"task-management-system-api/pkg/generators"
)

type contextKey string

const userCtxKey contextKey = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Ожидаем формат "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		// Проверяем токен
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return generators.GetJwtSecretKey(), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Добавляем userID в контекст запроса
		userID := claims.Subject
		ctx := context.WithValue(r.Context(), userCtxKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(userCtxKey).(string)
	if !ok {
		return "", fmt.Errorf("userID not found in context")
	}
	return userID, nil
}
