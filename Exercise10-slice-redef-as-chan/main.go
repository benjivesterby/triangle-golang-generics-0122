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
