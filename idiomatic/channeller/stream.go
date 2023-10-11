package channeller

import (
	"context"
	"time"
)

func StreamFn[I any, O any](ctx context.Context, in <-chan I, fn func(I) O) <-chan O {
	out := make(chan O)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v := <-in:
				out <- fn(v)
			default:
			}
		}
	}()

	return out
}

func StreamTimeoutFn[I any, O any](ctx context.Context, duration time.Duration, in <-chan I, fn func(I) O) (<-chan O, context.CancelFunc) {
	ctx2, cfn := context.WithTimeout(ctx, duration)
	return StreamFn(ctx2, in, fn), cfn
}
