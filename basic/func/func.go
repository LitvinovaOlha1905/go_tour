package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println(add(4, 7))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
