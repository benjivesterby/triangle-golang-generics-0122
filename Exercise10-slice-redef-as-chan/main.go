package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}

	fmt.Println("Printing data:")
	for v := range Slice[int](data).Chan() {
		fmt.Println(v)
	}

	data2 := Slice[int]{1, 2, 3, 4, 5}

	fmt.Println("Printing data 2:")
	for v := range data2.Chan() {
		fmt.Println(v)
	}
}

// TO USE: go get -u go.structs.dev/gen@latest
// COPIED FROM: https://github.com/structsdev/gen/blob/c58442d64fa19e214263757cf4d200de2eb2911c/slice.go#L30
type Slice[T any] []T

func (s Slice[T]) Chan() <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for _, v := range s {
			out <- v
		}
	}()

	return out
}
