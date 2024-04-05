package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Літерали структури
var (
	v1 = Vertex{1, 2}  // має тип Vertex
	v2 = Vertex{X: 1}  // Y:0 є неявним
	v3 = Vertex{}      // X:0 та Y:0
	p1  = &Vertex{1, 2} // має тип *Vertex
)

func main() {
	// Структура struct це набір полів.
	fmt.Println(Vertex{1, 2})

	// Доступ до полів структури здійснюється за допомогою крапки.
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Доступ до полів структури можна отримати через вказівник структури.
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// Літерал структури позначає щойно виділене значення структури шляхом переліку значень її полів.
	fmt.Println(v1, p1, v2, v3)
}
