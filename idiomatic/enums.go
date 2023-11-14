package idiomatic

//go:generate enumer -config=./caser/case-type.enum.yaml
//go:generate enumer -config=./math/geometry/angle-direction.enum.yaml
//go:generate enumer -config=./math/geometry/angle-type.enum.yaml
//go:generate enumer -config=./measurement/imperial-length-unit.enum.yaml
//go:generate enumer -config=./measurement/measurement-format-type.enum.yaml
//go:generate enumer -config=./measurement/measurement-system.enum.yaml
//go:generate enumer -config=./measurement/metric-unit.enum.yaml
//- go:generate enumer -path=./table.enum.go
