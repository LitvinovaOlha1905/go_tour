package main

import "fmt"

// fibonacci це функція, що повертає
// функцію, що повертає int.
func fibonacci() func() int {
a, b := 0, 1
    return func() int {
        result := a
        a, b = b, a+b
        return result
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}