package main

import "fmt"

func main() {

	// Форма range з циклу for дозволяє пройти (ітеруватись) по елементам зрізу чи мапи.
	// Під час застосування range до зрізу для кожної ітерації повертаються два значення. Перше — це індекс, а друге — копія елемента за цим індексом.
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Ви можете пропустити індекс або значення, вказавши _ замість змінної.
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}