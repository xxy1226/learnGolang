package main

import "fmt"

func main() {
	defer fmt.Println("world!") // defer 会立即求值，但是在外层函数结束后才执行

	fmt.Print("Hello ")
}
