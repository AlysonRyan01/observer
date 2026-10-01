package modelo

import "github.com/golang-jwt/jwt/v5"

type ClaimsUsuario struct {
	Role Role `json:"role"`
	jwt.RegisteredClaims
}
