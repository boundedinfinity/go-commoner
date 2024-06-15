package langer

var CSharp = cSharp{}

type cSharp struct{}

func (t cSharp) Identifier(s string) (string, error) {
	o := s

	return o, nil
}

func (t cSharp) MustIdentifier(s string) string {
	o, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return o
}
