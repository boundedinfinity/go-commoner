package measurement

//go:generate enumer -path=./measurement-format-type.enum.go

type MeasurmentFormat string

type measurmentFormats struct {
	Full         MeasurmentFormat `enum:"full"`
	Abbreviation MeasurmentFormat `enum:"abbreviation"`
}
