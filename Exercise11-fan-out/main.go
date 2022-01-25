package main

import "context"

func main() {
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
