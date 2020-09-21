package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int8
}

type Android struct {
	Person
	Model string
}

func (aPerson *Person) sayHi() {
	fmt.Printf("Hi, my name is %v, I'm %v years old.\n", aPerson.Name, aPerson.Age)
}

func (aDroid *Android) sayHello() {
	fmt.Printf("Hi, my name is %v, I'm %v years old, and my model name is %v.\n", aDroid.Name, aDroid.Age, aDroid.Model)
}
func main() {
	p1 := new(Person)
	p1.Name = "jack"
	p1.Age = 23
	p1.sayHi()

	p2 := new(Android)
	p2.Name = "little lee"
	p2.Age = 4
	p2.Model = "Nexus 5"
	p2.sayHello()
}
