package main

import "fmt"

// Довжина і ємність зрізу
func printSlice(s4 []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s4), cap(s4), s4)
}

// Створення зрізу за допомогою make
func printSlice2(s6 string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s6, len(x), cap(x), x)
}

// Додавання до зрізу
func printSlice3(s7 []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s7), cap(s7), s7)
}

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
		i  int
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

	// Довжина і ємність зрізу
	// Довжина зрізу (англ. length) - це кількість елементів, які він містить.
	// Ємність зрізу (англ. capacity) — це кількість елементів у базовому масиві, починаючи з першого елемента фрагмента.
	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)

	// Здійснити зріз зрізу, щоб встановити йому нульову довжину.
	s4 = s4[:0]
	printSlice(s4)

	// Збільшити його довжину.
	s4 = s4[:4]
	printSlice(s4)

	// Відкинути перші два значення.
	s4 = s4[2:]
	printSlice(s4)

	// Нульові зрізи

	var s5 []int
	fmt.Println(s5, len(s5), cap(s5))
	// if s5 == nil {
	// 	fmt.Println("nil!")
	// }

	// Створення зрізу за допомогою make
	// Функція make виділяє обнулений масив і повертає зріз, який посилається на цей масив
	a3 := make([]int, 5)
	printSlice2("a", a3)

	// Щоб указати ємність, передайте третій аргумент функції make

	b2 := make([]int, 0, 5)
	printSlice2("b", b2)

	c := b2[:3]
	printSlice2("c", c)

	d := c[1:5]
	printSlice2("d", d)

	// Додавання до зрізу
var s7 []int
	printSlice(s7)

	// append працює на 'nil' зрізах.
	s7 = append(s7, 0)
	printSlice3(s7)

	//  Зріз росте за потреби.
	s7 = append(s7, 1)
	printSlice3(s7)

	// Можна додавати більше одного елемента за раз.
	s7 = append(s7, 2, 3, 4)
	printSlice3(s7)

}
