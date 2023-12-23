package main

import "fmt"

type Digit int

func main() {
	const a = 5
	const b string = "Hello"

	fmt.Println(a, b)

	const (
		c = 5
		d = "Hello world"
	)

	fmt.Println(c, d)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		o Digit = iota
		_
		o2
		_
		o4
	)
	fmt.Println(o2, o4)
}
