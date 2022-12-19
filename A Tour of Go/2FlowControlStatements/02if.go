package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint((math.Sqrt(x)))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // 可以在 if 中运行一行代码
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // 这里不能使用 v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
