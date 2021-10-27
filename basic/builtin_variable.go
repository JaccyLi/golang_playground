package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func builtin() {
	fmt.Println("builtin")
}

func euler() {
	c := 3 + 4i
	fmt.Println("len of 3 + 4i-->", cmplx.Abs(c))
	fmt.Printf("pow(e, i * Pi)--> %3f\n", cmplx.Pow(math.E, 1i*math.Pi))
}

func forceTypeConvert() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func constble() {
	const fileLocation = "/home/suosuoli/"
	const a, b = 4, 5
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(fileLocation, c)
}

func enums() {

}

func main() {
	builtin()
	euler()
	forceTypeConvert()
	constble()
}
