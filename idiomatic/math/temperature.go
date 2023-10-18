package math

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToCelsius[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return (n - 32) * 5 / 9
	})
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToFahrenheit[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n*5/9 - 32
	})
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToKelvin[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToCelsius[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return n - 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToKelvin[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return (n-32)*5/9 + 273.15
	})
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToFahrenheit[T Numbers](x T) T {
	return singleToSingle(x, func(n float64) float64 {
		return (n-273.15)*5/9 - 32
	})
}
