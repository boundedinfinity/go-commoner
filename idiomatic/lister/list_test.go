package lister_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/lister"
	"github.com/stretchr/testify/assert"
)

func Test_List_PopB(t *testing.T) {
	s := lister.NewList[int]()
	s.PushB(1, 2)

	actual1, ok1 := s.PopB()
	assert.Equal(t, true, ok1)
	assert.Equal(t, 2, actual1)
	assert.Equal(t, 1, s.Len())

	actual2, ok2 := s.PopB()
	assert.Equal(t, true, ok2)
	assert.Equal(t, 1, actual2)
	assert.Equal(t, 0, s.Len())

	actual3, ok3 := s.PopB()
	assert.Equal(t, false, ok3)
	assert.Equal(t, 0, actual3)
	assert.Equal(t, 0, s.Len())
}

func Test_List_PopF(t *testing.T) {
	s := lister.NewList[int]()
	s.PushB(1, 2)

	actual1, ok1 := s.PopF()
	assert.Equal(t, true, ok1)
	assert.Equal(t, 1, actual1)
	assert.Equal(t, 1, s.Len())

	actual2, ok2 := s.PopF()
	assert.Equal(t, true, ok2)
	assert.Equal(t, 2, actual2)
	assert.Equal(t, 0, s.Len())

	actual3, ok3 := s.PopF()
	assert.Equal(t, false, ok3)
	assert.Equal(t, 0, actual3)
	assert.Equal(t, 0, s.Len())
}
