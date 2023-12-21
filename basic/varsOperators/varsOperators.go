package main

import "fmt"

func main() {
	// Определение переменных с типизацией
	var a int = 5          // 5
	var b string = "Hello" // Hello

	var c string // Пустая строка
	var d int    // 0

	fmt.Println(a, b, c, d)

	// Опеределение переменной сокращенным оператором присваивания
	e := 5
	f := "world"

	g, h := 5, 7
	i, h := 8, 6
	fmt.Println(e, f, g, i, h)

	i = 4
	fmt.Println(i)

	fmt.Println(i + 5)
	fmt.Println(i + h)
}
