package langer

var Java = golang{}

type java struct{}

func (t java) Identifier(s string) (string, error) {
	o := s

	return o, nil
}

func (t java) MustIdentifier(s string) string {
	o, err := t.Identifier(s)

	if err != nil {
		panic(err)
	}

	return o
}
