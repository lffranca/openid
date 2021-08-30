package jwt

import (
	"crypto/rsa"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type JWTParser struct{}

func (parser *JWTParser) Parser(tokenString string, key *rsa.PublicKey) (map[string]interface{}, error) {
	token, errToken := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errToken
}
