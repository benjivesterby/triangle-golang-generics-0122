package main

import "fmt"

func main() {
	data := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("Printing Keys:")
	for _, k := range Map[string, int](data).Keys() {
		fmt.Println(k)
	}

	fmt.Println("Printing Values:")
	for _, k := range Map[string, int](data).Values() {
		fmt.Println(k)
	}

	data2 := Map[string, int]{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("Printing data 2 Keys:")
	for _, k := range data2.Keys() {
		fmt.Println(k)
	}

	fmt.Println("Printing data 2 Values:")
	for _, k := range data2.Values() {
		fmt.Println(k)
	}
}

// TO USE: go get -u go.structs.dev/gen@latest
// COPIED FROM: https://github.com/structsdev/gen/blob/c58442d64fa19e214263757cf4d200de2eb2911c/map.go#L6
type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Keys() []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) Values() []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
