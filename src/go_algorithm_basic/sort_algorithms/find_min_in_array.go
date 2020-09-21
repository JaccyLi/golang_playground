package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 233, 63, 70,
		37, 2, 1, -22, 83, 27,
		19, 97, 999, 17,
	}
	var cnt int = 0
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if i != j {
				if x[i] < x[j] {
					cnt++
				}
			}
		}
		if cnt == len(x)-1 {
			fmt.Println("min=", x[i])
		}
		cnt = 0
	}
}
