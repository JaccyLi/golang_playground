package main

import "fmt"

func main() {
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)
	i := 1 == 1
	j := 1 == 2

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)

	var o bool
	fmt.Printf("%v, %T\n", o, o)


}
