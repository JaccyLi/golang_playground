package main

import (
	"fmt"
	"math"
)

type Animal struct {
	Genre  string
	Origin string
}

type Cat struct {
	Animal
	Name string
	Age  int8
	Host string
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// print 9 X 9 mutiplication chart
func main() {
	aAnimal := Animal{
		Genre:  "Vertebrate",
		Origin: "Sea monster",
	}
	aCat := Cat{}
	aCat.Animal.Genre = "mammal"
	fmt.Print(aAnimal)
	fmt.Print(aCat.Animal)

	c := Circle{2, 33, 5}
	fmt.Println(c.x)
	fmt.Println(c.y)
	fmt.Println(c.r)
	fmt.Println(c.area())
}
