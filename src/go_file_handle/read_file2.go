package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("go_file_handle/read.txt")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
