package idiomatic

//go:generate enumer -config=./caser/case-type.enum.yaml
//go:generate enumer -config=./mather/geometry/angle-direction.enum.yaml
//go:generate enumer -config=./mather/geometry/angle-type.enum.yaml
//go:generate enumer -config=./measurement/imperial-length-unit.enum.yaml
//go:generate enumer -config=./measurement/imperial-mass-unit.enum.yaml
//go:generate enumer -config=./measurement/measurement-format-type.enum.yaml
//go:generate enumer -config=./measurement/measurement-system.enum.yaml
//go:generate enumer -config=./measurement/metric-unit.enum.yaml
//go:generate enumer -config=./urier/uri-scheme.enum.yaml
//- go:generate enumer -path=./table.enum.go
