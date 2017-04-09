package xpkg

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Minus(x, y int) int {
	return x - y
}

func Desc() {
	fmt.Println("calculator")
}

func Add2(x, y int) int {
	return 2*x + 2*y
}
