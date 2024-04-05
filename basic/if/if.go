package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Оператор `if` з коротким виразом

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Оператори `if` та `else`
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	// Оператор `if` з коротким виразом

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Оператори `if` та `else`

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
