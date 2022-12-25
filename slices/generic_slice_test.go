package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenericSliceInt(t *testing.T) {
	glist := New[int]()
	glist.Insert(1)
	glist.Insert(2)
	glist.Insert(3)
	glist.Insert(4)

	require.Equal(t, 1, glist.Get(0))
	require.Equal(t, 2, glist.Get(1))
	require.Equal(t, 3, glist.Get(2))
	require.Equal(t, 4, glist.Get(3))

	glist.RemoveByIndex(2)
	require.Equal(t, 1, glist.Get(0))
	require.Equal(t, 2, glist.Get(1))
	require.Equal(t, 4, glist.Get(2))

	glist.RemoveByValue(2)
	require.Equal(t, 1, glist.Get(0))
	require.Equal(t, 4, glist.Get(1))
}

func Test_GenericSliceString(t *testing.T) {
	glist := New[string]()
	glist.Insert("1")
	glist.Insert("2")
	glist.Insert("3")
	glist.Insert("4")

	require.Equal(t, "1", glist.Get(0))
	require.Equal(t, "2", glist.Get(1))
	require.Equal(t, "3", glist.Get(2))
	require.Equal(t, "4", glist.Get(3))

	glist.RemoveByIndex(2)
	require.Equal(t, "1", glist.Get(0))
	require.Equal(t, "2", glist.Get(1))
	require.Equal(t, "4", glist.Get(2))

	glist.RemoveByValue("2")
	require.Equal(t, "1", glist.Get(0))
	require.Equal(t, "4", glist.Get(1))
}
