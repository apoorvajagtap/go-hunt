package common

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"ghunt/config"
	"fmt"
	"log"
	"reflect"
	"net/http"
	"strings"
	//"net/url"
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

func (c *CookieStruct) SetNewCookies() CookieStruct {

	var cookiesCollected CookieStruct
	val := reflect.ValueOf(cookiesCollected)
	typeOfCookies := val.Type()

	for i := 0; i<val.NumField(); i++ {
		
		if typeOfCookies.Field(i).Name == "CONSENT" {
			cookiesCollected.CONSENT = config.GetDefaultConsentCookie()
			continue
		}

		fmt.Print(typeOfCookies.Field(i).Name, " => ")
		var newValue string
		fmt.Scan(&newValue)

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
			break
		}
	}

 	return cookiesCollected
}

// Change the current working directory to allow using GHunt from anywhere
func GetSavedCookies() CookieStruct{
	// returns cookie cache if exists

	dataPath := config.GetDataPath()
	
	//dataPath := "/home/apjagtap/git-repos/personal-pr/go-hunt/data.txt"
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
	//fmt.Println("Detected stored cookies, checking it! \n", cookiesCollected)
	return content.Cookies

}

func GetAuthorizationSource(cookies CookieStruct) *http.Response{
	// returns html source of hangouts page if user authorized
	header := config.GetHeaders()
	
	req, _ := http.NewRequest("GET", "https://docs.google.com/document/u/0/?usp=direct_url", nil)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	
	val := reflect.ValueOf(cookies)
	typeOfCookies := val.Type()

	for i := 0; i<val.NumField(); i++ {
		req.AddCookie(&http.Cookie{Name: typeOfCookies.Field(i).Name, Value: val.Field(i).String()})
	}

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		req2, _ := http.NewRequest("GET", "https://hangouts.google.com", nil)
		for k, v := range header {
			req2.Header.Add(k, v)
		}
		for i := 0; i<val.NumField(); i++ {
			req2.AddCookie(&http.Cookie{Name: typeOfCookies.Field(i).Name, Value: val.Field(i).String()})
		}

		resp2, _ := client.Do(req2)
		defer resp2.Body.Close()

		data, _ := ioutil.ReadAll(resp2.Body)

		if strings.Contains(string(data), "myaccount.google.com") {
			return resp
		}
	}

	//fmt.Println(req, "#######  ", err, "###############")//, req.StatusCode)
	return nil
}

func CheckAndGen() {

	cookies_from_file := GetSavedCookies()
	var cookiesCollected CookieStruct
	new_cookies_entered := false

	if (cookies_from_file == CookieStruct{}) {

		fmt.Println("\nEnter these browser cookies found at accounts.google.com : ")
		cookiesCollected = cookies_from_file.SetNewCookies()

		fmt.Println("lalalala >>>", cookiesCollected)
	} else {
	// in case user wants to enter new cookies (example: for new account)
		html := GetAuthorizationSource(cookies_from_file)
		valid := false

		if html != nil {
			fmt.Println("\n[+] The cookies seems valid !")
			valid = true
		} else {
			fmt.Println("\n[-] Seems like the cookies are invalid.")
		}

		fmt.Print("\nDo you want to enter new browser cookies from accounts.google.com ? (Y/n) ")
		var new_gen_inp string
		fmt.Scan(&new_gen_inp)

		if strings.ToLower(new_gen_inp) == "y" {
			new_cookies_entered = true
			cookiesCollected := cookies_from_file.SetNewCookies()
			fmt.Println(new_cookies_entered, cookiesCollected)
		} else if valid == false {
			log.Fatal("Please put valid cookies. Exiting...")
		}
	}

	// Valid Cookies
	if (new_cookies_entered == true || cookies_from_file == CookieStruct{}) {
		html := GetAuthorizationSource(cookiesCollected)
		if html != nil {
			fmt.Println("\n[+] The cookies seems valid !")
		} else {
			log.Fatal("\n[-] Seems like the cookies are invalid, try regenerating them.")
		}
	}

	if new_cookies_entered != true {
		cookiesCollected = cookies_from_file
		fmt.Print("Do you want to generate new tokens ? (Y/n) ")
		var choice string
		fmt.Scan(&choice)

		if strings.ToLower(choice) == "y" {
			os.Exit(0)
		}
	}


	// Start the exctraction Process

	// We first initialize the browser driver
	//chrome_options := GetChromeOptionsArgs()

	fmt.Println("chcchchchc >>> ", cookies_from_file)
}