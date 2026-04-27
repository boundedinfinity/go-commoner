package temperature

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToCelsius[T mather.Number](x T) T {
	return (x - 32) * 5 / 9
}

// https://byjus.com/temperidiomatic.Numbersion-formula/
func CelsiusToFahrenheit[T mather.Number](x T) T {
	return x*5/9 - 32
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToKelvin[T mather.Number](x T) T {
	return T(float64(x) + 273.15)
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToCelsius[T mather.Number](x T) T {
	return T(float64(x) - 273.15)
}

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToKelvin[T mather.Number](x T) T {
	return T((float64(x)-32)*5/9 + 273.15)
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToFahrenheit[T mather.Number](x T) T {
	return T(((float64(x)-273.15)*5/9 - 32))
}
