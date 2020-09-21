package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(suffix string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	cl := makeSuffix(".jpg")
	fmt.Printf("name : %v\n", cl("name"))
	fmt.Printf("kid.jpg : %v\n", cl("kid.jpg"))
	fmt.Printf("girl : %v\n", cl("girl"))
}
