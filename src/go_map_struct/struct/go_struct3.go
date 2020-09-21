package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func totalArea(s ...Shape) float64 {
	var a float64
	for _, aas := range s {
		a += aas.area()
	}
	return a
}
func totalPerimeter(p ...Shape) float64 {
	var pp float64
	for _, pa := range p {
		pp += pa.perimeter()
	}
	return pp
}
func main() {
	c1 := Circle{4, 5, 10}
	r1 := Rectangle{0, 0, 10, 10}
	fmt.Printf("area of c1 is %v\nperimeter of c1 is %v\narea of r1 is %v\nperimeter of r1 is %v\n",
		c1.area(), c1.perimeter(), r1.area(), r1.perimeter())
	fmt.Printf("total area of c1 and r1 is: %v\ntotal perimeter of c1 and r1 is: %v",
		totalArea(&c1, &r1), totalPerimeter(&c1, &r1))
}
