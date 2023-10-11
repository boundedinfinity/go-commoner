package rational

import (
	"fmt"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"golang.org/x/exp/constraints"
)

func IntegerComponent[T constraints.Float](x T) int {
	i, _ := math.Modf[T, int](x)
	return i
}

func DecimalComponent[T constraints.Float](n T) int {
	_, f := math.Modf[T, float64](n)
	str := fmt.Sprintf("%e", f)
	str = stringer.Split(str, "e-")[1]
	pow, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	sizef := f * math.Pow10(pow+1)
	sizei := int(math.Round(sizef))
	return sizei
}
