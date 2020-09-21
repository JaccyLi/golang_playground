package main

import (
	"fmt"
	"time"
)

func main() {
	goto l2
l1:
	fmt.Println("I'm l1, going to l2 ...")
	time.Sleep(1000 * time.Millisecond)
l2:
	fmt.Println("I'm l2, going to l1 ...")
	time.Sleep(1000 * time.Millisecond)
	goto l1
}
