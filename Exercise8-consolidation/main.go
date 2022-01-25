package main

type MyStruct struct {
	Data int
	Tags string
}

func main() {

	mystringchan := make(chan string)
	myintchan := make(chan int)
	mystructchan := make(chan MyStruct)

	SToIChan(mystringchan)
	IntToIChan(myintchan)
	StrToIChan(mystructchan)

	ToAnyChan(mystringchan)
	ToAnyChan(myintchan)
	ToAnyChan(mystructchan)
}

// Example of channel conversions without generics
func SToIChan(in <-chan string) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}

func IntToIChan(in <-chan int) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}

func StrToIChan(in <-chan MyStruct) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}

// Example of channel conversion using generics
func ToAnyChan[T any](in <-chan T) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		for s := range in {
			out <- s
		}
	}()
	return out
}
