package main

import "fmt"

func main() {
	s1 := []bool{true, false, true, true} // 不指定长度的是切片，可以扩容
	fmt.Println(s1)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:]
	fmt.Println(s)
}
