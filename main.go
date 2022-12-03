package main

import "fmt"

func dosometh(a, b int, A func(int, int) int) int {

	return A(a, b)
}

func main() {
	fmt.Println(dosometh(10, 5, func(x, y int) int { return x + y }))
	fmt.Println(dosometh(10, 5, func(x, y int) int { return x / y }))
	fmt.Println(dosometh(10, 5, func(x, y int) int { return x * y }))
}
