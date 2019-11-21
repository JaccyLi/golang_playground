package main

import "fmt"


func main() {
	a := 10
	b := 5
	var c int = 20
	var d int16 = 50

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	
	fmt.Println(b << 3)
	fmt.Println(b >> 3)

	fmt.Println(a % b)
	fmt.Println(c + int(d))
}
