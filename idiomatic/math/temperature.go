package math

import "golang.org/x/exp/constraints"

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToCelsius[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return (n - 32) * 5 / 9
	})
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToFahrenheit[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n*5/9 - 32
	})
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToKelvin[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToCelsius[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n - 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToKelvin[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return (n-32)*5/9 + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToFahrenheit[T constraints.Float | constraints.Integer](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return (n-273.15)*5/9 - 32
	})
}
