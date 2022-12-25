package slices

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SortingByInc(t *testing.T) {
	numbers := Numbers{1, 10, 4, 9, 3}
	sort.Sort(byInc{numbers})

	expected := Numbers{1, 3, 4, 9, 10}

	require.Equal(t, expected, numbers)
}

func Test_SortingByDec(t *testing.T) {
	numbers := Numbers{1, 10, 4, 9, 3}
	sort.Sort(byDec{numbers})

	expected := Numbers{10, 9, 4, 3, 1}

	require.Equal(t, expected, numbers)
}
