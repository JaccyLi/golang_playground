package main

import (
	"os"
)

func main() {
	file, err := os.Create("go_file_handle/create.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Hello, I'm created by os.\n")
}
