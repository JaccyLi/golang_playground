package utils

import "fmt"

type myInt int
type myFunType func(int, int) int

func sum(m int, n int) int {
	return m + n
}

func sumF1(funVar func(int, int) int, o int, p int) int {
	return funVar(o, p)
}

func sumF2(funVar myFunType, o int, p int) int {
	return funVar(o, p)
}

func sub(m int, n int) {

}
func run() {
	var i myInt
	i = 40
	fmt.Printf("i = %v\n type i : %T\n", i, i)
	fa := sum
	k := sum(2, 3)
	//kk := fa(33, 44)
	f1 := sumF1(sum, 12, 23)
	f2 := sumF2(sum, 22, 13)

	fmt.Printf("type fa is: %T\nk = %v\nf1 = %v\nf2 = %v\n",
		fa, k, f1, f2)
}
