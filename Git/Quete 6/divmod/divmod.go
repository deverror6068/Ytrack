package main

import "fmt"

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	Divmod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

func Divmod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}