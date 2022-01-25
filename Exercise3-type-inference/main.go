package main

func main() {
	// Non-Inferred Call
	Inferred[int](1)

	// Inferred Call
	Inferred(1)
}

func Inferred[T any](in T) T {
	return in
}
