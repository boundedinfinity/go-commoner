package langer

// https://go.dev/ref/spec#Identifiers

var Go = golang{}

type golang struct{}

func (t golang) Identifier(s string) (string, error) {
	o := s

	return o, nil
}

func (t golang) MustIdentifier(s string) string {
	o, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return o
}
