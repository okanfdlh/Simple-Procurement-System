package config

import "os"

func JwtSecret() string {
	return os.Getenv("JWT_SECRET")
}
