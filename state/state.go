package state

import (
	"encoding/json"
	"os"
	"time"
)

const validityPeriod = time.Minute * 5

type State struct {
	Value   string `json:"value"`
	Expires int64  `json:"expires"`
}

func EncryptState(value string) (string, error) {
	state := State{Value: value, Expires: time.Now().Add(validityPeriod).Unix()}
	jsonData, err := json.Marshal(state)
	if err != nil {
		return "", err
	}
	return Encrypt(string(jsonData), os.Getenv("STATE_PASSWORD"))
}

func DecryptState(value string) (state State, err error) {
	data, err := Decrypt(value, os.Getenv("STATE_PASSWORD"))
	if err != nil {
		return State{}, err
	}

	if err = json.Unmarshal([]byte(data), &state); err != nil {
		return State{}, err
	}
	return state, nil
}
