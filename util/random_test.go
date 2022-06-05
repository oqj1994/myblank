package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	s := RandomString(5)
	require.NotEmpty(t, s)
	require.Equal(t, 5, len(s))
}

func TestRandomNum(t *testing.T) {
	n := RandomInt(5, 10)
	require.NotEmpty(t, n)
	require.Less(t, n, int64(10))
	require.Greater(t, n, int64(5))
}
