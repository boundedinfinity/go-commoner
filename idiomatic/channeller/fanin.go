package channeller

import "sync"

func FanIn[T any](inputs ...<-chan T) chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	merge := func(c <-chan T) {
		defer wg.Done()

		for elem := range c {
			out <- elem
		}
	}

	wg.Add(len(inputs))

	for _, c := range inputs {
		go merge(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
