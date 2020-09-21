package utils

import "fmt"

var Age int
var Name string

func init() {
	Age = 20
	Name = "jack"
	fmt.Printf("Age = %v\n Name = %v\n", Age, Name)
}

