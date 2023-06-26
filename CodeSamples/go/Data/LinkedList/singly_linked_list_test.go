package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseSuccess(t *testing.T) {
	l := New(1, 2, 3, 4, 5)

	require.Equal(t, "[1,2,3,4,5]", l.String())

	l.Reverse()
	require.Equal(t, "[5,4,3,2,1]", l.String())
}

func TestReversedReverse(t *testing.T) {
	l := New(1, 2, 3, 4, 5)

	l.Reverse().Reverse()
	require.Equal(t, "[1,2,3,4,5]", l.String())
}
