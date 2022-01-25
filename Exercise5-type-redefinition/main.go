package main

import "fmt"

func main() {

	var x MyType = 1

	fmt.Println(BadFunc(1))
	// BadFunc(x)

	fmt.Println(GoodFunc(1))
	fmt.Println(GoodFunc(x))
}

type MyType int

func BadFunc[T int](in T) T {
	return in
}

func GoodFunc[T ~int](in T) T {
	return in
}
