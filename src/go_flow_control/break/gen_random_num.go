package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var cnt int
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		cnt++
		fmt.Println("n = ", n)
		if n == 99 {
			break
		}
	}
	fmt.Println(cnt)
}
