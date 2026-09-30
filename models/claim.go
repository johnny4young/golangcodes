package models

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
