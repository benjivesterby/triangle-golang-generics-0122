package main

import (
	"constraints"
	"fmt"
	"math/rand"
	"time"
)

// Initialize the random number generator.
func init() { rand.Seed(time.Now().Unix()) }

func main() {
	fmt.Printf("Random Int8: %v\n", Int[int8]())
	fmt.Printf("Random UInt8: %v\n", Int[uint8]())
	fmt.Printf("Random Int16: %v\n", Int[int16]())
	fmt.Printf("Random UInt26: %v\n", Int[uint16]())
	fmt.Printf("Random Int32: %v\n", Int[int32]())
	fmt.Printf("Random UInt32: %v\n", Int[uint32]())
	fmt.Printf("Random Int: %v\n", Int[int]())
	fmt.Printf("Random UInt: %v\n", Int[uint]())
	fmt.Printf("Random Float32: %v\n", Float[float32]())
	fmt.Printf("Random Float64: %v\n", Float[float64]())

}

func Int[T constraints.Integer]() T {
	value := rand.Int()
	return T(value)
}

func Float[T constraints.Float]() T {
	return T(rand.Float64())
}

func Ints[T constraints.Integer](size int) []T {
	out := make([]T, size)

	for i := range out {
		out[i] = Int[T]()
	}

	return out
}

func Floats[T constraints.Float](size int) []T {
	out := make([]T, size)

	for i := range out {
		out[i] = Float[T]()
	}

	return out
}

func IntTests[T constraints.Integer](tests, cap int) [][]T {
	out := make([][]T, tests)

	for i := range out {
		out[i] = Ints[T](cap)
	}

	return out
}

func FloatTests[T constraints.Float](tests, cap int) [][]T {
	out := make([][]T, tests)

	for i := range out {
		out[i] = Floats[T](cap)
	}

	return out
}
