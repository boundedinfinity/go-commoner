package langer

// https://mathiasbynens.be/notes/javascript-identifiers

var JavaScript = golang{}

type javaScript struct{}

func (t javaScript) Identifier(s string) (string, error) {
	o := s

	return o, nil
}

func (t javaScript) MustIdentifier(s string) string {
	o, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return o
}
