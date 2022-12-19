package main

import "fmt"

// 尽量在短函数中这么用。在长函数中会因为找不到返回值而影响可读性
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(10))
}
