package main

import "context"

func main() {
}

// TO USE: go get -u go.atomizer.io/stream@latest
// COPIED FROM: https://github.com/devnw/stream/blob/8186e9bc20761e99c3906122c054ea0ce09d7a31/stream.go#L32
type InterceptFunc[T, U any] func(context.Context, T) (U, bool)

// TO USE: go get -u go.atomizer.io/stream@latest
// COPIED FROM: https://github.com/devnw/stream/blob/8186e9bc20761e99c3906122c054ea0ce09d7a31/stream.go#L32
// Intercept accepts an incoming data channel and a function literal that
// accepts the incoming data and returns data of the same type and a boolean
// indicating whether the data should be forwarded to the output channel.
// The function is executed for each data item in the incoming channel as long
// as the context is not cancelled or the incoming channel remains open.
func Intercept[T, U any](
	ctx context.Context,
	in <-chan T,
	fn InterceptFunc[T, U],
) <-chan U {
	if ctx == nil {
		ctx = context.Background()
	}

	out := make(chan U)

	go func() {
		defer recover()
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				// Executing this in a function literal ensures that any panic
				// will be caught during execution of the function
				func() {
					// TODO: Should something happen with this panic data?
					defer recover()

					// Determine if the function was successful
					result, ok := fn(ctx, v)
					if !ok {
						return
					}

					// Execute the function against the incoming value
					// and send the result to the output channel.
					select {
					case <-ctx.Done():
						return
					case out <- result:
					}
				}()
			}
		}
	}()

	return out
}
