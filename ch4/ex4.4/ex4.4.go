package main

import "fmt"

func main(){
	a := 3
	var b float64 = 3.5

	var c int = int(b) // 3
	d := float64(a * c) // 9.0

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3)
	var h int = int(b) * 3
	fmt.Println(g, h ,f)

}