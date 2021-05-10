package common

import (
	"os"
	//"io/ioutil"
	"encoding/json"
	"ghunt/config"
	"fmt"
	"log"
	"reflect"
)

type Config struct {
	Hangouts string `json:"hangouts_auth"`
	Internal string `json:"internal_auth"`
	Keys struct {
		GDoc string `json:"gdoc"`
		HangOK string `json:"hangouts"`
		IntK string `json:"internal"`
	} `json:"keys"`
	Cookies CookieStruct `json:"cookies"`
}

type CookieStruct struct {
	SID string `json:"sid"`
	SSID string `json:"ssid"`
	APISID string `json:"apisid"`
	SAPISID	string `json:"sapisid"`
	HSID string `json:"hsid"`
	CONSENT string `json:"consent"`
}

func (c *CookieStruct) setCookies(cookiesCollected CookieStruct) CookieStruct {
	//var cook CookieStruct
	val := reflect.ValueOf(cookiesCollected)
	typeOfCookies := val.Type()

	for i := 0; i<val.NumField(); i++ {
		//fName := typeOfCookies.Field(i).Name
		
		fmt.Println(typeOfCookies.Field(i).Name, "=>")
		var newValue string
		fmt.Scanln(&newValue)

		switch {
		case typeOfCookies.Field(i).Name == "SID":
			cookiesCollected.SID = newValue
			break
		case typeOfCookies.Field(i).Name == "SSID":
			cookiesCollected.SSID = newValue
			break
		case typeOfCookies.Field(i).Name == "APISID":
			cookiesCollected.APISID = newValue
			break
		case typeOfCookies.Field(i).Name == "SAPISID":
			cookiesCollected.SAPISID = newValue
			break
		case typeOfCookies.Field(i).Name == "HSID":
			cookiesCollected.HSID = newValue
			break
		default:
			cookiesCollected.CONSENT = config.GetDefaultConsentCookie()
		}
	}

// 	fmt.Println("\nEnter these browser cookies found at accounts.google.com : \n SID => ")
// 	var sid_new string
// 	fmt.Scanln(&sid_new)
// 	cook.SID = sid_new

// 	fmt.Println("SSID => ")
// 	var ssid_new string
// 	fmt.Scanln(&ssid_new)
// 	cook.SSID = ssid_new

// 	fmt.Println("APISID => ")
// 	var api_new string
// 	fmt.Scanln(&api_new)
// 	cook.APISID = api_new

// 	fmt.Println("SAPISID => ")
// 	var sapi_new string
// 	fmt.Scanln(&sapi_new)
// 	cook.SAPISID = sapi_new

// 	fmt.Println("HSID => ")
// 	var hsid_new string
// 	fmt.Scanln(&hsid_new)
// 	cook.HSID = hsid_new

// 	cook.CONSENT = config.GetDefaultConsentCookie()

 	return cookiesCollected
}

// Change the current working directory to allow using GHunt from anywhere
func GetSavedCookies() {
	// returns cookie cache if exists

	//dataPath := config.GetDataPath()
	
	dataPath := "/home/apjagtap/git-repos/personal-pr/go-hunt/data.txt"
	fmt.Println("here?", )
	if _, err := os.Stat(dataPath); err != nil{
		log.Fatal(err)
	}

	var content Config
	configFile, err := os.Open(dataPath)
	defer configFile.Close()
	if err != nil{
		log.Fatal(err)
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&content)
	cookiesCollected := content.Cookies
	fmt.Println("Detected stored cookies, checking it! \n", cookiesCollected)

	//new_cookies_entered := false

	if (cookiesCollected == CookieStruct{}) {
		fmt.Println("BEFOREEEEE >>>>>>>>>>>\n", cookiesCollected)
		
		//c_return := cookiesCollected.setCookies()

		fmt.Println("\nEnter these browser cookies found at accounts.google.com : ")
		c_return := cookiesCollected.setCookies(cookiesCollected)
		// val := reflect.ValueOf(cookiesCollected)
		// typeOfCookies := val.Type()

		// for i := 0; i<val.NumField(); i++ {
		// 	fName := typeOfCookies.Field(i).Name
			
		// 	fmt.Println(typeOfCookies.Field(i).Name, "=>")
		// 	var newValue string
		// 	fmt.Scanln(&newValue)
		// }
	

		fmt.Println(c_return)
	}

	

}