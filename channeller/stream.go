package channeller

func StreamFn[I any, O any](in <-chan I, fn func(I) O) <-chan O {
	out := make(chan O)

	go func() {
		for item := range in {
			out <- fn(item)
		}

		close(out)
	}()

	return out
}
