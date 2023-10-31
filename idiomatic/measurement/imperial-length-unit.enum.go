package measurement

// https://en.wikipedia.org/wiki/Imperial_units

//go:generate enumer -path=./imperial-length-unit.enum.go

type ImperialLengthUnit string

type imperialLengthUnits struct {
	Twip         ImperialLengthUnit `enum:"twip"`
	Thou         ImperialLengthUnit `enum:"thow"`
	Barleycorn   ImperialLengthUnit `enum:"barleycorn"`
	Inch         ImperialLengthUnit `enum:"inch"`
	Hand         ImperialLengthUnit `enum:"hand"`
	Foot         ImperialLengthUnit `enum:"foot"`
	Yard         ImperialLengthUnit `enum:"yard"`
	Chain        ImperialLengthUnit `enum:"chain"`
	Furlong      ImperialLengthUnit `enum:"furlong"`
	Mile         ImperialLengthUnit `enum:"mile"`
	League       ImperialLengthUnit `enum:"league"`
	Fathom       ImperialLengthUnit `enum:"fathom"`
	Cable        ImperialLengthUnit `enum:"cable"`
	NauticalMile ImperialLengthUnit `enum:"nautical-mile"`
	Link         ImperialLengthUnit `enum:"link"`
	Rod          ImperialLengthUnit `enum:"rod"`
	cf           map[ImperialLengthUnit]float64
	abbr         map[ImperialLengthUnit][]string
}
