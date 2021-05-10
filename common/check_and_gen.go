package common

import (
	"os"
	//"io/ioutil"
	"encoding/json"
	"ghunt/config"
	"fmt"
	"log"
)

type CookieCollect struct {
	Hangouts string `json:"hangouts_auth"`
	Internal string `json:"internal_auth"`
	Keys struct {
		GDoc string `json:"gdoc"`
		HangOK string `json:"hangouts"`
		IntK string `json:"internal"`
	} `json:"keys"`
	Cookies struct {
		SID string `json:"sid"`
		SSID string `json:"ssid"`
		APISID string `json:"apisid"`
		SAPISID	string `json:"sapisid"`
		HSID string `json:"hsid"`
		CONSENT string `json:"consent"`
	} `json:"cookies"`
}

// Change the current working directory to allow using GHunt from anywhere
func GetSavedCookies() {
	// returns cookie cache if exists
	dataPath := config.GetDataPath()
	fmt.Println("here?", )
	if _, err := os.Stat(dataPath); err == nil{

		var content CookieCollect
		configFile, err := os.Open(dataPath)
		defer configFile.Close()
		if err != nil{
			log.Fatal(err)
		}

		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&content)

		fmt.Println("check now!!", content.Keys.GDoc)
	}
}