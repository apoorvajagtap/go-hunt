package main

import (
	"fmt"
	"os"
	"strings"
	"ghunt/helpers"
	"ghunt/common"
	// "github.com/apoorvajagtap/go-hunt/modules/doc_hunt"
	// "github.com/apoorvajagtap/go-hunt/modules/email_hunt"
)

func strExists(slice []string, sub_string string) bool {
	for _, ele := range slice {
		if ele == sub_string {
			return true
		}
	}

	return false
}

func main() {
	modules := []string{"doc", "email"}

	if len(os.Args) <= 1 || !strExists(modules, strings.ToLower(os.Args[1])) {
		fmt.Println("Please choose a module.")
		fmt.Println("Available modules :")

		for _, module := range modules {
			fmt.Println(module)
		}
		os.Exit(0)
	}

	module := strings.ToLower(os.Args[1])
	//fmt.Println(module)
	// if len(os.Args) >= 3 {
	// 	data := os.Args[2]
	// }

	fmt.Println("calling check&gen")
	//common.GetSavedCookies()
	common.CheckAndGen()

	data := ""
	//data := "https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	//data := "https://docs.google.com/document/d/edit#123"


	if module == "email" {
		fmt.Println(helpers.DeclareEmail())
		//helpers.EmailHunt(data)

	} else if module == "doc" { 
		//fmt.Println(helpers.DeclareDoc())
		helpers.DocHunt(data)
	}

}

// STORY LINE
// to compare the string, first used sort.SearchStrings.find() method, but as it panics if element is not found.
// We defined custom func to check through each element of slice and return a bool value
