package keystorsa

import (
	"log"

	"github.com/lffranca/openid/internal/entities"
)

func NewkeysToRSA(parser JWTParser, key *entities.Key) *keysToRSA {
	if parser == nil {
		log.Fatalln("the \"JWTParser\" cannot be null")
	}

	if key == nil {
		log.Fatalln("the \"key\" cannot be null")
	}

	return &keysToRSA{
		Parser: parser,
		Key:    key,
	}
}

type keysToRSA struct {
	Parser JWTParser
	Key    *entities.Key
}

func (item *keysToRSA) Validate(token string) (map[string]interface{}, error) {
	key, errKey := item.Key.ToRSA()
	if errKey != nil {
		return nil, errKey
	}

	return item.Parser.Parser(token, key)
}
