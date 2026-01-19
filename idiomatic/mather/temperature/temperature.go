package temperature

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/internal"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
)

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToCelsius[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return (n - 32) * 5 / 9
	})
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToFahrenheit[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n*5/9 - 32
	})
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToKelvin[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToCelsius[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n - 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToKelvin[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return (n-32)*5/9 + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToFahrenheit[T types.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return (n-273.15)*5/9 - 32
	})
}
