package langer

// https://mathiasbynens.be/notes/javascript-identifiers

var TypeScript = typeScript{}

type typeScript struct{}

func (t typeScript) Identifier(s string) (string, error) {
	o := s

	return o, nil
}

func (t typeScript) MustIdentifier(s string) string {
	o, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return o
}
