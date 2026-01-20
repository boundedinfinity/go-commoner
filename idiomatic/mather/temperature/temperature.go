package temperature

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/internal"
)

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToCelsius[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return (n - 32) * 5 / 9
	})
}

// https://byjus.com/temperidiomatic.Numbersion-formula/
func CelsiusToFahrenheit[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n*5/9 - 32
	})
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToKelvin[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToCelsius[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return n - 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToKelvin[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return (n-32)*5/9 + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToFahrenheit[T idiomatic.Number](x T) T {
	return internal.SingleToSingle[T, T](x, func(n float64) float64 {
		return (n-273.15)*5/9 - 32
	})
}
