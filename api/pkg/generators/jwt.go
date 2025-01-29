package generators

import (
	"os"
)

func getJwtAccessKey() []byte {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return []byte("secret_key")
	}

	return []byte(secretKey)
}
