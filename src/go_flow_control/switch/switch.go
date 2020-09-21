package main

import (
	"fmt"
)

func isGolang(attr string) bool {
	var a bool
	switch attr {
	case "multiMethods":
		a = false
	case "concise":
		a = true
	default:
		a = false
	}
	return a
}

func main() {
	//is := isGolang("concise")
	//is := isGolang("multiMethods")
	is := isGolang("")
	fmt.Printf("he is a gopher? %v", is)
}
