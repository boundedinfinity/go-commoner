package pipeline

func New() *Pipeline {
	return &Pipeline{}
}

type Pipeline struct {
}

// type Mapper[I any, O any] interface {
// 	Run(I) (O, error)
// }

// // type MapFn func[I any, O any](I) (O, error)
// // type WalkFunc func(path string, info fs.FileInfo, err error) error

// func (t *Pipeline) Map(mapper Mapper) {

// }
