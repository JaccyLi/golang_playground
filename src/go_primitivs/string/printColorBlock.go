package main

import (
	"fmt"
)

const (
	InfoColor    = "\033[1;44m \033[0m"
	NoticeColor  = "\033[1;46m \033[0m"
	WarningColor = "\033[1;43m \033[0m"
	ErrorColor   = "\033[1;41m \033[0m"
	BackColor    = "\033[0;45m \033[0m"
)

var (
	colBlanks int
	rowBlanks string
)

// move cursor to (cols, rows) position
func move(col int, row int) {
	for i := 0; i < row; i++ {
		fmt.Println("")
	}
	for j := 0; j < col; j++ {
		rowBlanks += " "
	}
	fmt.Printf(rowBlanks)
}

// cursor position (x(cols), y(rows)), block size (m, n)
func printBlock(x int, y int, m int, n int, color string) {
	//var blockLen string
	move(x, y)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			fmt.Print(color)
		}
		fmt.Println("")

		rowBlanks = ""
		for k := 0; k < x; k++ {
			rowBlanks += " "
		}
		fmt.Print(rowBlanks)
	}
	rowBlanks = ""
	fmt.Println("")
}

func main() {

	//cols, rows := consolesize.GetConsoleSize()
	//fmt.Printf("%v %v", cols, rows)
	printBlock(10, 0, 10, 5, InfoColor)
	printBlock(10, 0, 10, 5, WarningColor)
	printBlock(10, 0, 10, 5, ErrorColor)
	printBlock(20, 0, 10, 5, WarningColor)
	printBlock(30, 0, 10, 5, ErrorColor)

}
