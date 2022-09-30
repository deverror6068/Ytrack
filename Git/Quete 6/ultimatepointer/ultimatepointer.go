package main

import "fmt"

func PointOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	n := &b
	PointOne(&n)
	fmt.Println(a)
}
