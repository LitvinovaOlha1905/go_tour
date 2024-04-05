package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Привіт"
	a[1] = "Світ"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices

	var s []int = primes[1:4]
	fmt.Println(s)

	// Зміна елементів зрізу змінює відповідні елементи його основного масиву.
	// Інші зрізи, які мають той самий базовий масив, побачать ці зміни.

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a2 := names[0:2]
	b := names[1:3]
	fmt.Println(a2, b)

	b[0] = "XXX"
	fmt.Println(a2, b)
	fmt.Println(names)

	// Літерали зрізів
	// Літерал зрізу схожий на літерал масиву без довжини.

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b2 bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	// Зріз за замовчуванням
	// Під час створення зрізу ви можете не зазначати верхню або нижню межі, щоб замість них використовувати стандартні значення.
	// За замовчуванням для нижньої межі встановлено нуль, а для верхньої – довжина зрізу.

	s3 := []int{2, 3, 5, 7, 11, 13}

	s3 = s3[1:6]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)

	s3 = s3[:]
	fmt.Println(s3)
}