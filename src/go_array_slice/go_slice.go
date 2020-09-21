package main

import "fmt"

func main() {
	slice1 := make([]string, 123, 333)
	fmt.Printf("type: %T, value %v, len %v, cap %v", slice1, slice1, len(slice1), cap(slice1))
}
