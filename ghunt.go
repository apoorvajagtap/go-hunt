package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/apoorvajagtap/go-hunt/modules/doc_hunt"
	"github.com/apoorvajagtap/go-hunt/modules/email_hunt"
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

	module = strings.ToLower(os.Args[1])
	if len(os.Args) >= 3{
		data = os.Args[2]
	}
	else{
		data = None
	}

	if module == "email"{
		email_hunt(data)
	}
	else if module == "doc"{
		doc_hunt(data)
	}
}

// STORY LINE
// to compare the string, first used sort.SearchStrings.find() method, but as it panics if element is not found.
// We defined custom func to check through each element of slice and return a bool value
