package main

import (
	"fmt"
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

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(abc, abcd, str, stri)
}
