package models

import "github.com/dgrijalva/jwt-go"

// Dog struct for dog object.
type Dog struct {
	Owner  string `json:"owner"`
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
}

// User struct for user object.
type User struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Token struct for token object.
type Token struct {
	Login string `json:"login" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// Bird struct for bird object
type Bird struct {
	Owner string `json:"owner" bson:"owner"`
	Name  string `json:"name" bson:"name" validate:"required"`
	Type  string `json:"type" bson:"type" validate:"required"`
}

// CustomClaims struct for jwt token generation claim.
type CustomClaims struct {
	Login string
	jwt.StandardClaims
}
