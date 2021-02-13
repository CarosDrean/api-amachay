package models

import jwt "github.com/dgrijalva/jwt-go"

type Claim struct {
	ClaimUser `json:"user"`
	jwt.StandardClaims
}


