package channeller

func FanOut[T any](input <-chan T, outputs ...<-chan T) []chan T {
	var newOutputs []chan T

	return newOutputs
}
