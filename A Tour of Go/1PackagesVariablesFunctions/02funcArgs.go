package main

import "fmt"

func add(x int, y int) int { // Can be shorted as:	func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
