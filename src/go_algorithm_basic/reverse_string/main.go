package main

import "fmt"

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[len(s)-1-i], s[i] = s[i], s[len(s)-1-i]
	}
	fmt.Printf("%c", s)
}

func main() {
	var ss = []byte{'A', ' ', 'm', 'a', 'n', ',', ' ', 'a', ' ', 'p', 'l', 'a', 'n', ',', ' ', 'a', ' ', 'c', 'a', 'n', 'a', 'l', ':', ' ', 'P', 'a', 'n', 'a', 'm', 'a'}
	fmt.Println("len(ss)==", len(ss))
	fmt.Printf("unreverse: %c\n", ss)
	reverseString(ss)
}
