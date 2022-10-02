package case_type

//go:generate enumer -path=main.go

type CaseType string

const (
	Camel      CaseType = "camel"
	Kebab      CaseType = "kebab"
	Kebabupper CaseType = "kebabUpper"
	Pascal     CaseType = "pascal"
	Phrase     CaseType = "phrase"
	Snake      CaseType = "snake"
	Snakeupper CaseType = "snakeUpper"
	Unknown    CaseType = "unknown"
)
