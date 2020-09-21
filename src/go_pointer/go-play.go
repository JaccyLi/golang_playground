package main

import (
	"fmt"
	_ "math"
	_ "unsafe"
)

var (
	num1 int     = 8
	num2 float64 = 3.1415926
)

func main() {

	var ptr1 *int
	var ptr2 *float64
	ptr1 = &num1
	ptr2 = &num2
	fmt.Printf("num1=%v\nnum2=%v\nptr1=%v\nptr2=%v\nptr1_value=%v\nptr2_value=%v\n",
		num1, num2, ptr1, ptr2, *ptr1, *ptr2)

	var i = 222
	fmt.Println("b ", i)
	mo(&i)
	fmt.Println("a ", i)

	k := 1.5
	var r = square(&k)
	fmt.Println(r)

	var x, y = 33, 32432
	var m, n = swap(&x, &y)
	fmt.Println(m, n)
}

func mo(ptr *int) {
	*ptr = 666
}

func square(k *float64) float64 {
	*k = *k * *k
	return *k
}

func swap(x *int, y *int) (int, int) {
	*x, *y = *y, *x
	return *x, *y
}
