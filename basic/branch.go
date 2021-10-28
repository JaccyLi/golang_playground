package main

import (
	"fmt"
	"os"
)

func readF(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", f)
	}
}

func readF2(file string) {
	if f, err := os.ReadFile(file); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", f)
	}
}

func score(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d\n", score))
	case score == 100:
		g = "A"
	case score > 90:
		g = "B"
	case score > 80:
		g = "C"
	case score > 70:
		g = "D"
	}
	return g
}

func main() {
	const file = "./read.file"
	readF(file)
	readF2(file)

	fmt.Println(
		//score(-1),
		score(74),
		score(90),
		score(95),
		score(100),
		score(101),
	)

}
