package utils

import (
	"fmt"
	"strconv"
)

/*

ABBC

AB + BC = AC + BB

*/

func lvWan(a, b, c int) {
	ab := strconv.Itoa(a) + strconv.Itoa(b)
	bc := strconv.Itoa(b) + strconv.Itoa(c)
	ac := strconv.Itoa(a) + strconv.Itoa(c)
	bb := strconv.Itoa(b) + strconv.Itoa(b)
	intab, _ := strconv.Atoi(ab)
	intbc, _ := strconv.Atoi(bc)
	intac, _ := strconv.Atoi(ac)
	intbb, _ := strconv.Atoi(bb)
	if intab+intbc == intac+intbb {
		fmt.Println("Got: ", ab+bc)
	}
	//fmt.Printf("ab: %v\nbc: %v\nac: %v\nbb: %v", ab, bc, ac, bb)
}
func get() {
	var P int = 20
	for i := 0; i <= P; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k <= j; k++ {
				if j+1 == i && k+1 == j {
					lvWan(k, j, i)
				}
			}
		}
	}
}
