package main

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

func GetCircuitUrl() (string, string, string, string) {
	btcSoPath := os.Getenv("btcSoPath")
	btcSetup := os.Getenv("btcSetup")
	ethSoPath := os.Getenv("ethSoPath")
	ethSetup := os.Getenv("ethSetup")
	return btcSoPath, btcSetup, ethSoPath, ethSetup
}
