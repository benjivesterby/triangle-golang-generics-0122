package main

import "context"

func main() {
}

// COPIED FROM:https://github.com/devnw/stream/blob/8186e9bc20761e99c3906122c054ea0ce09d7a31/stream.go#L87
// FanIn accepts incoming data channels and forwards returns a single channel
// that receives all the data from the supplied channels.
//
// NOTE: The transfer takes place in a goroutine for each channel
// so ensuring that the context is cancelled or the incoming channels
// are closed is important to ensure that the goroutine is terminated.
func FanIn[T any](ctx context.Context, in ...<-chan T) <-chan T {
	if ctx == nil {
		ctx = context.Background()
	}

	out := make(chan T)

	if len(in) == 0 {
		defer close(out)
		return out
	}

	defer func() {
		go func() {
			<-ctx.Done()
			close(out)
		}()
	}()

	for _, i := range in {
		// Pipe the result of the channel to the output channel.
		go Pipe(ctx, i, out)
	}

	return out
}

// COPIED FROM: https://github.com/devnw/stream/blob/8186e9bc20761e99c3906122c054ea0ce09d7a31/stream.go#L19
// Pipe accepts an incoming data channel and pipes it to the supplied
// outgoing data channel.
//
// NOTE: Execute the Pipe function in a goroutine if parallel execution is
// desired. Cancelling the context or closing the incoming channel is important
// to ensure that the goroutine is properly terminated.
func Pipe[T any](ctx context.Context, in <-chan T, out chan<- T) {
	if ctx == nil {
		ctx = context.Background()
	}
	// Pipe is just a fan-out of a single channel.
	FanOut(ctx, in, out)
}

// COPIED FROM: https://github.com/devnw/stream/blob/8186e9bc20761e99c3906122c054ea0ce09d7a31/stream.go#L117
// FanOut accepts an incoming data channel and copies the data to each of the
// supplied outgoing data channels.
//
// NOTE: Execute the FanOut function in a goroutine if parallel execution is
// desired. Cancelling the context or closing the incoming channel is important
// to ensure that the goroutine is properly terminated.
func FanOut[T any](ctx context.Context, in <-chan T, out ...chan<- T) {
	if ctx == nil {
		ctx = context.Background()
	}

	if len(out) == 0 {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}

			for _, o := range out {
				// Closure to catch panic on closed channel write.
				// Continue Loop
				func() {
					defer recover()
					select {
					case <-ctx.Done():
						return
					case o <- v:
					}
				}()
			}
		}

	}
}
