package main

import (
	"constraints"
	"fmt"
	"math/rand"
)

type MyType int

func main() {
	var x []MyType = []MyType{1, 3, 4, 3, 4, 2, 234, 523, 45, 234524, 2345}

	fmt.Printf("Len of ints: %d\n", Len([]int{1, 2, 3}))
	fmt.Printf("Len of mytype: %d\n", Len(x))
	fmt.Printf("Len of byte slice: %d\n", Len([]byte{1, 2, 3}))
	fmt.Printf("Len of string: %d\n", Len([]string{"hello", "world"}))

	for _, v := range Ints[int](100) {
		fmt.Printf("%d %% 121 == %d\n", v, ModInt(v))
	}

}

type MyAwesomeConstraint[T comparable] interface {
	~[]T
}

func Len[U MyAwesomeConstraint[T], T comparable](in U) int {
	return len(in)
}

type Inty interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func ModInt[T Inty](in T) T {
	return in % 121
}

func Int[T constraints.Integer]() T {
	value := rand.Int()
	return T(value)
}

func Ints[T constraints.Integer](size int) []T {
	out := make([]T, size)

	for i := range out {
		out[i] = Int[T]()
	}

	return out
}
