package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Літерали мап
// Літерали мапи подібні до структурних літералів, але з обов'язковим використанням ключів.

var m2 = map[string]Vertex{

	// "Bell Labs": Vertex{
	//	40.68433, -74.39967,
	//},
	//"Google": Vertex{
	//	37.42202, -122.08408,
	//},

	// Якщо тип верхнього рівня є просто ім'ям типу, то його можна не включати в елементи літералу.
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {

	// Функція make повертає мапу заданого типу, ініціалізовану та готову до використання.
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m2)

	// // Зміна мап
	m3 := make(map[string]int)

	// //Вставка або оновлення елементу в мапі m: m[key] = elem
	m3["Відповідь"] = 42
	fmt.Println("Значення:", m3["Відповідь"])

	// // Отримати елемент: elem = m[key]
	m3["Відповідь"] = 48
	fmt.Println("Значення:", m3["Відповідь"])

	// // Видалити елемент: delete(m, key)
	delete(m3, "Відповідь")
	fmt.Println("Значення:", m["Відповідь"])

	// // Перевірити наявність ключа з допомогою присвоєння двох значень: elem, ok = m[key]
	v, ok := m3["Відповідь"]
	fmt.Println("Значення:", v, "Присутнє?", ok)
	// Якщо key є в m, то ok рівне true. Якщо ні, то ok рівне false.
	// Якщо key відсутній у мапі, то elem є нульовим значенням для типу елементу мапи.
	// Якщо elem або ok ще не були оголошені, можна використати коротку форму оголошення:
	// elem, ok := m[key]

}
