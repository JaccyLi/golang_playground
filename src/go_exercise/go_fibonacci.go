package main

import "fmt"

func fib(n int) {
	var a int = 0
	var b int = 1
	for i := 0; i < n; i++ {
		fmt.Printf(" %v ", a)
		var fa int = a + b
		a = b
		b = fa
	}
}
func main() {
	fib(20)
}
