package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RemoveAndNotKeepOrder(t *testing.T) {
	s := MySlice{1, 3, 2, 4, 5}
	s = s.Remove(2)
	if len(s) != 4 {
		t.Errorf("incorrect length of slice; expected 4 but got %d", len(s))
	}
	require.Equal(t, MySlice{1, 3, 5, 4}, s)
}

func Test_RemoveAndKeepOrder(t *testing.T) {
	s := MySlice{1, 33, 10, 4, 5}
	s = s.RemoveAndKeepOrder(2)
	if len(s) != 4 {
		t.Errorf("incorrect length of slice; expected 4 but got %d", len(s))
	}
	require.Equal(t, MySlice{1, 33, 4, 5}, s)
}
