package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	// 有条件的 switch ，对比条件与 case 的值决定
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	// 没有条件的 switch ，根据 case 的布尔值决定
	switch {
	case t.Hour() < 12:
		fmt.Println(("Good morning!"))
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
