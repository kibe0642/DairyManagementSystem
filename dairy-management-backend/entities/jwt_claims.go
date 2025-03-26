package entities

import "github.com/golang-jwt/jwt/v4"

// JWTClaims represents the structure of JWT claims
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
