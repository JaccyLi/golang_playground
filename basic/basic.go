package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var abc = 123
var str = "suosuoli"

// abc := 123 // invalid

// short way
var (
	abcd = 1234
	stri = "suosuolili"
)

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var c string = "hello"
	fmt.Println(a, b, c)
}

func variableTypeDeduction() {
	var d, e, f = 2, true, "hello"
	fmt.Println(d, e, f)
}

func variableShort() {
	d, e, f := 2, true, "hello"
	d = 3
	fmt.Println(d, e, f)
}

func builtin() {
	fmt.Println("builtin-----")
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
	fmt.Println("enums----------")
	const (
		apple  = 0
		orange = 1
		banana = 2
	)

	const (
		c = iota
		cpp
		_
		golang
		javascript
		fortran
		scalar
		ruby
		shell
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(apple, orange, banana, c, cpp, golang, javascript, scalar, ruby, shell)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(abc, abcd, str, stri)
	builtin()
	euler()
	forceTypeConvert()
	constble()
	enums()
}
