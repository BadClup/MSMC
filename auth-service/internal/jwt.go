package internal

import (
	"github.com/dgrijalva/jwt-go"
	"math"
	"msmc/auth-service/shared"
)

type jwtPayload struct {
	Exp      int64  `json:"exp"`
	Email    string `json:"email"`
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
}

func signJwt(payload jwtPayload) (string, error) {
	if payload.Exp == 0 {
		payload.Exp = math.MaxInt64
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      payload.Exp,
		"email":    payload.Email,
		"username": payload.Username,
		"user_id":  payload.UserID,
	})

	return token.SignedString([]byte(shared.JwtSecret))
}
