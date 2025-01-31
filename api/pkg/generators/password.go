package generators

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func GenerateHashPassword(password string) string {
	salt := os.Getenv("SALT")
	if salt == "" {
		salt = "SALT"
	}

	hash := md5.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}
