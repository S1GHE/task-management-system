package generators

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

func getJwtSecretKey() []byte {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return []byte("secret_key")
	}

	return []byte(secretKey)
}

func generateAccessToken(userId uuid.UUID) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userId.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJwtSecretKey())
}

func generateRefreshToken() (string, error) {
	randomBytes := make([]byte, 32)

	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes), nil
}
