package main

import "fmt"

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println(c())

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
