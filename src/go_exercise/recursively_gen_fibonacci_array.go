package main

import "fmt"

// fib(n) = fib(n-1) + fib(n-2)
// 0 1 1 2 3 5 8 13 ...
func main() {
	genfib(12)
}

func genfib(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(fib(uint(i)))
	}
}

func fib(n uint) uint {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
