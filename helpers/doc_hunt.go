package helpers

import(
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"net/url"
	validator "github.com/asaskevich/govalidator"
)

func DeclareDoc() string {
	return "This one's DOC !!"
}


func DocHunt(doc_path string) {
	
	urlValid := validator.IsURL(doc_path)
	if urlValid != true {
		fmt.Println("Please enter a valid URL..")
		os.Exit(1)
	}
	
	filtered_url, _ := url.Parse(doc_path)
	filtered_url.Fragment = ""

	//fmt.Println("HAHAHAHAHHAHAHA  >>> ", filtered_url)

	split_doc_url := strings.Split(strings.Split(doc_path, "?")[0], "/")
	var doc_id string

	for _, ele := range split_doc_url {
		if len(ele) == 44 {
			doc_id = ele
			break
		}
	}
	
	if doc_id != "" {
		fmt.Println("Document ID: ", doc_id)
	} else {
		fmt.Println("\nDocument ID not found.\nPlease make sure you have something that looks like this in your link :1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms")
		os.Exit(1)
	}
	//doc_id := split_doc_url[len(split_doc_url)-1]
	//fmt.Println(doc_path, "[[[[[[[[[[[", doc_path_1, "\\\\\\\\\\\\", doc_path_list)

	// DO NOT REMOVE THE FOLLOWING PART !!!!
	// if _, err := os.Stat("config.data_path"); err != nil {
	// 	fmt.Println("Please generate cookies and tokens first, with the check_and_gen.py script.")
	// 	os.Exit(1)
	// }

	internal_token := ""
	cookies := {}


	content, err := ioutil.ReadFile("config.data_path")
	//file, err := os.Open("config.data_path")
}