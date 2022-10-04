package lister_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/lister"
	"github.com/stretchr/testify/assert"
)

func Test_List_PopB(t *testing.T) {
	s := lister.NewList[int]()
	s.PushB(1, 2)

	actual := s.PopB()
	assert.Equal(t, true, actual.Defined())
	assert.Equal(t, 2, actual.Get())
	assert.Equal(t, 1, s.Len())

	actual = s.PopB()
	assert.Equal(t, true, actual.Defined())
	assert.Equal(t, 1, actual.Get())
	assert.Equal(t, 0, s.Len())

	actual = s.PopB()
	assert.Equal(t, false, actual.Defined())
	assert.Equal(t, 0, actual.Get())
	assert.Equal(t, 0, s.Len())
}

func Test_List_PopF(t *testing.T) {
	s := lister.NewList[int]()
	s.PushB(1, 2)

	actual := s.PopF()
	assert.Equal(t, true, actual.Defined())
	assert.Equal(t, 1, actual.Get())
	assert.Equal(t, 1, s.Len())

	actual = s.PopF()
	assert.Equal(t, true, actual.Defined())
	assert.Equal(t, 2, actual.Get())
	assert.Equal(t, 0, s.Len())

	actual = s.PopF()
	assert.Equal(t, false, actual.Defined())
	assert.Equal(t, 0, actual.Get())
	assert.Equal(t, 0, s.Len())
}
