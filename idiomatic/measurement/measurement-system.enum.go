package measurement

//go:generate enumer -path=./measurement-system.enum.go

type MeasurementSystem string

type measurementSystems struct {
	Metric   MeasurementSystem `enum:"metric"`
	Imperial MeasurementSystem `enum:"imperial"`
}
