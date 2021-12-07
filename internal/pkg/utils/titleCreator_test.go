package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTitleCreator(t *testing.T) {
	req := require.New(t)
	cases := func(i int, str, want string) func(t *testing.T) {
		return func(t *testing.T) {
			result := TitleCreator(i, str)
			req.Equal(result, want)
		}
	}

	t.Run("Success case #1", cases(1, "123", "1_123.md"))
	t.Run("Success case #2", cases(2, "/123", "2_123.md"))
	t.Run("Success case #3", cases(3, "/1/2/3 test words", "3_123_test_words.md"))
	t.Run("Success case #4", cases(4, "/1/2/3 test words______", "4_123_test_words.md"))
}

func TestStringDeleteExtraChar(t *testing.T) {
	req := require.New(t)
	cases := func(str, want string) func(t *testing.T) {
		return func(t *testing.T) {
			result := stringDeleteExtraChar(str)
			req.Equal(result, want)
		}
	}

	t.Run("Success case #1", cases("123", "123"))
	t.Run("Success case #2", cases("/123", "123"))
	t.Run("Success case #3", cases("/1/2/3", "123"))
	t.Run("Success case #4", cases("/1/2/3/", "123"))
	t.Run("Success case #5", cases("-1.2@3#", "123"))
	t.Run("Success case #6", cases("-1.2@3#_____", "123"))
}
