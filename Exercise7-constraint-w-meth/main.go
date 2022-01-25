package main

import "fmt"

type MyType int

func (m MyType) String() string {
	return "MyType"
}

func main() {

	// BadFunc(1)
	// GoodFunc(1)
	fmt.Println(GoodFunc(MyType(1)))
}

type BadInty interface {
	int | int8 | int16 | int32 | int64
	String() string
}

func BadFunc[T BadInty](in T) T {
	return in
}

type Inty interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
}

func GoodFunc[T Inty](in T) T {
	return in
}
