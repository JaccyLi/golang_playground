package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("read.txt")
	if err != nil {
		// handle the error here
		fmt.Println("file not found!")
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("file read error!")
		return
	}

	str := string(bs)
	fmt.Println(str)
}
