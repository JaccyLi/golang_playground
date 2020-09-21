package main

import (
	"fmt"
	ut "go_function/utils"
)

/*
1. 先执行引入包的变量定义

2. 执行引入包的init函数

3. 执行main所在的包的变量定义

4. 执行main所在的包的init函数

5. 执行main函数
*/

var PI = 3.1415926
var tvar = tt()

func tt() int {
	fmt.Println("tt() ...")
	return 20
}

func init() {
	fmt.Println("init() ...")
}

func main() {

	fmt.Printf("main() ... age = %v\n main() ... name = %v\n", ut.Age, ut.Name)
}
