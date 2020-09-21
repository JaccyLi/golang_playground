package main

import (
	"fmt"
)

// print 9 X 9 mutiplication chart
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v X %v = %v\t", j, i, j*i)
		}
		fmt.Println("")
	}
}
