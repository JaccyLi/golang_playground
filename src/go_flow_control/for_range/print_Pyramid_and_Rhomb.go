package main

import "fmt"

// print a pyramid
/*
       *        1
	  ***		3
	 *****		5
	*******     7
   *********    9
  ***********   11
 *************  13
*************** 15
       *
	  * *     2 * floors - 1
	 *   *
	*     *
   *       *
  *         *
 *           *
***************
|
|
|
       *
	  * *
	 *   *
	*     *
   *       *
  *         *
 *           *
*             *
 *           *
  *         *
   *       *
	*     *
	 *   *
	  * *
       *

*/

func main() {
	printPyramid(6)
	printRhomb(6)
	printVacuumRhomb(6)
}

func printPyramid(floor int) {
	for i := 1; i <= floor; i++ {
		for m := 1; m <= floor-i; m++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == floor {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func printRhomb(floor int) {
	for i := 1; i <= 2*floor-1; i++ {
		// split floors to two parts
		if i <= floor {
			// print blanks before print *
			for m := 1; m <= floor-i; m++ {
				fmt.Print(" ")
			}
			// print *
			for m := 1; m <= 2*i-1; m++ {
				fmt.Print("*")
			}
		} else {
			for n := 1; n <= i-floor; n++ {
				fmt.Print(" ")
			}
			for n := 2*((2*floor-1)-i) + 1; n >= 1; n-- {
				fmt.Print("*")
			}
		}
		// give a new line
		fmt.Println("")
	}
}

func printVacuumRhomb(floor int) {
	for i := 1; i <= 2*floor-1; i++ {
		// split floors to two parts
		if i <= floor {
			// print blanks before print *
			for m := 1; m <= floor-i; m++ {
				fmt.Print(" ")
			}

			// print *
			for m := 1; m <= 2*i-1; m++ {
				if m == 1 || m == 2*i-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}

		} else {
			for n := 1; n <= i-floor; n++ {
				fmt.Print(" ")
			}
			for n := 2*((2*floor-1)-i) + 1; n >= 1; n-- {
				if n == 1 || n == 2*((2*floor-1)-i)+1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		// give a new line
		fmt.Println("")
	}

}
