package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // &i 为指向 i 的指针
	fmt.Println(*p) // *p 为指针 p 所指向目标的值
	*p = 21         // 将 指针 p 指向的目标值赋为21
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
