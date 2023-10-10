package caser

//go:generate enumer -path=./case-type.enum.go

type CaseType string

type caseTypes struct {
	Camel      CaseType
	Kebab      CaseType
	Kebabupper CaseType
	Pascal     CaseType
	Phrase     CaseType
	Snake      CaseType
	Snakeupper CaseType
	Unknown    CaseType
}
