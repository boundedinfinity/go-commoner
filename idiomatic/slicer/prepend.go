package slicer

// References
//      - https://www.scala-lang.org/api/3.4.2/scala/collection/ArrayOps.html#prependedAll-fffff738

func Prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}
