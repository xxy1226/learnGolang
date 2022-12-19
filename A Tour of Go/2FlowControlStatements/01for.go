package main

import "fmt"

func main() {
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	sum2 := 1
	for sum2 < 1000 { // 这里的 for 就是 while
		sum2 += sum2
	}
	fmt.Println(sum2)
}
