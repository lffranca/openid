package keystorsa

import "crypto/rsa"

type JWTParser interface {
	Parser(token string, key *rsa.PublicKey) (claims map[string]interface{}, err error)
}
