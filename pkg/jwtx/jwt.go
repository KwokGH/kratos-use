package jwtx

import "github.com/golang-jwt/jwt/v5"

func CreateJwtToken(secretKey string, iat int64, seconds int64, metadata map[string]interface{}) (string, error) {
	claims := jwt.MapClaims(metadata)
	claims["exp"] = iat + seconds
	claims["iat"] = iat

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
