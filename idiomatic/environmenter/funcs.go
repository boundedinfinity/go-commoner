package environmenter

func SubstitueOk(elem string) (string, bool) {
	return New().SubstitueOk(elem)
}

func Substitue(elem string) string {
	return New().Substitue(elem)
}

func SubAll(elems ...string) []string {
	return New().SubstitueAll(elems...)
}
