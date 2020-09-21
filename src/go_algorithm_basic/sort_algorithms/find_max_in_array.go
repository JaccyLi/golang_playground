package main

import "fmt"

func main() {
	var a = []int{-4441, 123, 222, 55, 9999, 2, 3}
	v := findMax(a...)
	fmt.Print(v)
}

func findMax(a ...int) int {
	var cnt int = 0
	var max int
	for i, vi := range a {
		for j, vj := range a {
			if i != j {
				if vi > vj {
					cnt++
				}
			}
		}
		if cnt == len(a)-1 {
			max = vi
		}
		cnt = 0
	}
	return max
}
