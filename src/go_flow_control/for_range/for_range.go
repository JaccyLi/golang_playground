package main

import "fmt"

type stu struct {
	name     string
	age      int8
	location string
	friends  []string
}

var quot = "北京"

func main() {
	stu1 := stu{
		name:     "jack",
		age:      24,
		location: "Beijing",
		friends:  []string{"john", "bob", "mick"},
	}

	for _, ch := range stu1.friends {
		fmt.Println(ch)
	}

	for i, ch := range quot {
		fmt.Printf("%d %#U\n", i, ch)
	}
	fmt.Println(stu1)
}
