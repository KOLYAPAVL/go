package main

import "fmt"

func main() {
	var a int8 = -128
	var b int16 = -500
	var c int32 = -55555555
	var d int64 = -2348132734323
	fmt.Println(a, b, c, d)

	var e uint8 = 128
	var f int16 = 257
	var g int32 = 1234231
	var h int64 = 5000000000
	fmt.Println(e, f, g, h)

	var i int = -5000
	var j uint = 20000
	fmt.Println(i, j)

	var k float32 = 3.6
	var l float64 = 3.8888
	fmt.Println(k, l)

}
