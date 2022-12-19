package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v)

	p := &v
	p.X = 1e9 // 隐式简介引用 explicit dereference
	fmt.Println(v)
}
