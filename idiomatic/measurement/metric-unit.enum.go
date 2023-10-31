package measurement

//go:generate enumer -path=./metric-unit.enum.go

type MetricUnit string

type metricUnits struct {
	Tera  MetricUnit `enum:"tera"`
	Giga  MetricUnit `enum:"giga"`
	Mega  MetricUnit `enum:"mega"`
	Kilo  MetricUnit `enum:"kilo"`
	Hecto MetricUnit `enum:"hecto"`
	Deca  MetricUnit `enum:"deca"`
	Unit  MetricUnit `enum:"unit"`
	Deci  MetricUnit `enum:"deci"`
	Centi MetricUnit `enum:"centi"`
	Milli MetricUnit `enum:"milli"`
	Micro MetricUnit `enum:"micro"`
	Nano  MetricUnit `enum:"nano"`
	Pico  MetricUnit `enum:"pico"`
	cf    map[MetricUnit]float64
	abbr  map[MetricUnit][]string
}
