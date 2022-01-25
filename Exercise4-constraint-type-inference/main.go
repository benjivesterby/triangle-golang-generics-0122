package main

import (
	"fmt"
	"math/rand"
)

func main() {
	out := ConstraintTypeInference(1)
	fmt.Printf("%T\n", out)

	out2 := ConstraintTypeInference("hello world")
	fmt.Printf("%T\n", out2)
}

// ConstraintTypeInference demonstrates that type inference works for
// U which is a composite of T.
func ConstraintTypeInference[U []T, T any](in T) U {
	count := rand.Int() % 100
	out := make([]T, count)

	for i := 0; i < count; i++ {
		out[i] = in
	}

	return out
}
