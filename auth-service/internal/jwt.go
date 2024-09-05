package internal

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"math"
	"msmc/auth-service/shared"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type jwtPayload struct {
	Exp      int64  `json:"ExpiresAt"`
	Email    string `json:"email"`
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
}

func signJwt(payload jwtPayload) (string, error) {
	if payload.Exp == 0 {
		payload.Exp = math.MaxInt32
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": payload.Exp,
		"email":     payload.Email,
		"username":  payload.Username,
		"user_id":   payload.UserID,
	})

	return token.SignedString([]byte(shared.JwtSecret))
}

func decodeJwt(tokenString string) (jwtPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(shared.JwtSecret), nil
	})
	if err != nil {
		return jwtPayload{}, err
	}

	if !token.Valid {
		return jwtPayload{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwtPayload{}, errors.New("invalid token claims")
	}

	return jwtPayload{
		Exp:      int64(claims["ExpiresAt"].(float64)),
		Email:    claims["email"].(string),
		Username: claims["username"].(string),
		UserID:   int(claims["user_id"].(float64)),
	}, nil
}
