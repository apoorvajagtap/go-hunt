package common

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"ghunt/config"
	"fmt"
)

type JSONResponse struct{
	Hangouts string `json:"hangouts_auth"`
	Internal string `json:"internal_auth"`
	Keys Nested `json:"nested_keys"`
	Cookies_ID Cookies `json:"cookies"` 
}

type Nested struct {
	Doc string `json:"google_doc"`
	HangoutPath string `json:"hangouts"`
	InternalPath string `json:"internal"`
}

type Cookies struct{
	SID string `json:"sid_cookie"`
	SSID string `json:"ssid_cookie"`
	APISID string `json:"apisid_cookie"`
	SAPISID string `json:"sapsid_cookie"`
	HSID string `json:"hsid_cookie"`
	CONSENT string `json:"consent"`
}

// Change the current working directory to allow using GHunt from anywhere
func GetSavedCookies() {
	// returns cookie cache if exists
	dataPath := config.GetDataPath()
	fmt.Println("here?", )
	if _, err := os.Stat(dataPath); err == nil{
		//file, err := os.Open(data_path)
		//defer file.Close()

		byteValue, _ := ioutil.ReadFile(dataPath)
		var result map[string]string

		json.Unmarshal([]byte(byteValue), &result)

		fmt.Println("reached here!!", result)

	}
}