package keystorsa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/lffranca/openid/internal/entities"
	"github.com/lffranca/openid/pkg/jwt"
)

func TestKeysToRSA_Converter(t *testing.T) {
	url := os.Getenv("DISCOVERY_KEYS_URL")
	log.Println(url)
	resp, errResp := http.Get(url)
	if errResp != nil {
		t.Error(errResp)
		return
	}

	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		t.Error(errBody)
		return
	}

	var configs struct {
		Keys []*entities.Key `json:"keys"`
	}
	if err := json.Unmarshal(body, &configs); err != nil {
		t.Error(err)
		return
	}

	if len(configs.Keys) <= 0 {
		t.Errorf("not valid keys")
		return
	}

	parser := &jwt.JWTParser{}

	tokenByte, errToken := ioutil.ReadFile(os.Getenv("JWT_FILE_PATH"))
	if errToken != nil {
		t.Error(errToken)
		return
	}

	usecase := NewkeysToRSA(parser, configs.Keys[0])
	claims, err := usecase.Validate(string(tokenByte))
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(claims)
}
