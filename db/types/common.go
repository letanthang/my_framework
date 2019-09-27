package types

import "github.com/dgrijalva/jwt-go"

type ErrorResponse struct {
	Code    string
	Message string
}

type MyClaims struct {
	jwt.StandardClaims
	ClientID    int    `json:"cid"`
	ClientType  string `json:"typ"`
	ClientName  string `json:"noc"`
	ClientPhone string `json:"poc"`
	ClientEmail string `json:"eoc"`
}
