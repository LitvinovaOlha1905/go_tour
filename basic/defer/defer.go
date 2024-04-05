package main

import "fmt"

func main() {
	// Оператор defer відкладає виконання функції до тих пір, поки оточуюча функція не повертає значення.
	defer fmt.Println("світ")
	fmt.Println("Привіт")

	// Складування defer
	// Виклики відкладених функцій додаються на стек. При поверненні з функції, відкладені виклики виконуються у порядку "останній прийшов - перший пішов".

	fmt.Println("Рахуємо")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Завершено")
}
