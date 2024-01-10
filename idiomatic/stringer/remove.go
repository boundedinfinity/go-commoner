package stringer

func Remove[T ~string, U ~string](s T, remove U) string {
	return Replace[T, U, string](s, remove, "")
}

func RemoveNumbers[T ~string](s T) string {
	return ReplaceNumbers(s, "")
}

func RemoveLetters[T ~string](s T) string {
	return ReplaceLetters(s, "")
}

func RemoveSymbols[T ~string](s T) string {
	return ReplaceSymbols(s, "")
}

func RemoveNonNumbers[T ~string](s T) string {
	return ReplaceNonNumbers(s, "")
}

func RemoveNonLetters[T ~string](s T) string {
	return ReplaceNonLetters(s, "")
}

func RemoveNonWord[T ~string](s T) string {
	return ReplaceNonWord(s, "")
}

func RemoveSpace[T ~string](s T) string {
	return ReplaceSpace(s, "")
}

func RemoveSpaces[T ~string](s T) string {
	return ReplaceSpaces(s, "")
}

func RemoveNewlines[T ~string](s T) string {
	return ReplaceNewlines(s, "")
}
