package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestKeyToRSA(t *testing.T) {
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
		Keys []*Key `json:"keys"`
	}
	if err := json.Unmarshal(body, &configs); err != nil {
		t.Error(err)
		return
	}

	pathToSave := os.Getenv("JWT_FILE_PATH")
	for index, key := range configs.Keys {
		pemKey, _ := key.ToBytes()
		path := fmt.Sprintf(pathToSave, key.Kid)
		log.Printf("Key %d: %s\n", index, path)
		if err := ioutil.WriteFile(path, pemKey, 0644); err != nil {
			log.Println("Error save file: ", err)
		}
	}
}
