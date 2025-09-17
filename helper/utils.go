package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func GetBtcUrl() (string, string, string) {
	btcUrl := os.Getenv("btcUrl")
	btcPwd := os.Getenv("btcPwd")
	btcUser := os.Getenv("btcUser")
	return btcUrl, btcUser, btcPwd
}

func GetCircuitUrl() (string, string) {
	btcSetup := os.Getenv("btcSetup")
	ethSetup := os.Getenv("ethSetup")
	return btcSetup, ethSetup
}

func GetSoPath() (string, string) {
	btcSoPath := os.Getenv("btcSoPath")
	ethSoPath := os.Getenv("ethSoPath")
	return btcSoPath, ethSoPath
}
func ToObj(s string, obj interface{}) error {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}
	return json.Unmarshal([]byte(s), obj)
}

func ToJson(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
