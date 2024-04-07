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

// Закриття функцій (англ. function closures)
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println(add(4, 7))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	// Закриття функцій (англ. function closures)
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
