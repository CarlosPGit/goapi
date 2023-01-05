package jwt_config

import (
	"time"

	"github.com/CarlosPGit/golang_sface/internal/admin/user"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

const expTime = 30

func NewJWT() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	})
}

func GenerateToken(user user.User) (string, error) {
	claims := jwt.MapClaims{
		"name":  user.Email,
		"admin": true,
		"exp":   time.Now().Add(time.Minute * expTime).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
}
