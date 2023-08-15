package main

import "fmt"

func main() {

	lang := make(map[string]string)
	lang["js"] = "javascript"
	lang["rc"] = "replication controller"

	fmt.Println(lang)
	fmt.Println(lang["js"])

	delete(lang, "rc")
	fmt.Println(lang)

	//loops
	lang["ps"] = "jkdfhkds"
	for _, val := range lang {
		fmt.Println(val)
	}
}
