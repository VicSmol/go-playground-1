package stringsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	t.Run("should normalize string", func(t *testing.T) {
		var inputs = []string{"", "   ", " b ", "hello, world  ", "  hello,   world  "}
		var expected = []string{"", "", "b", "hello, world", "hello, world"}
		var result = make([]string, len(inputs))

		for index, input := range inputs {
			result[index] = Normalize(input)
		}

		assert.Equal(t, expected, result)
	})
}

func TestSplit(t *testing.T) {
	t.Run("should split string", func(t *testing.T) {
		input := []string{"", ",,,,,,", "a", ",,a, ,b,,c,,", ",,a=b,,c=d,e=f,,"}
		expected := [][]string{{}, {}, {"a"}, {"a", " ", "b", "c"}, {"a=b", "c=d", "e=f"}}
		result := make([][]string, len(input))
		separator := ","

		for index, input := range input {
			result[index] = Split(input, separator)
		}

		assert.Equal(t, expected, result)
	})
}

func TestJoin(t *testing.T) {
	t.Run("should join slice", func(t *testing.T) {
		input := [][]string{
			{},
			{"a"},
			{"a", "b", "c"},
			{"aa", "bb", "cc"},
		}
		expected := []string{"", "a", "a,b,c", "aa,bb,cc"}
		result := make([]string, len(input))
		separator := ","

		for index, str := range input {
			result[index] = Join(str, separator)
		}

		assert.Equal(t, expected, result)
	})
}

func TestParseKV(t *testing.T) {
	t.Run("should parse string", func(t *testing.T) {
		input := []string{"", "a=b;c=d", ";;a=b;;c=d;e=f;;"}
		expected := []map[string]string{
			{},
			{"a": "b", "c": "d"},
			{"a": "b", "c": "d", "e": "f"},
		}
		result := make([]map[string]string, len(input))

		for index, str := range input {
			result[index] = ParseKV(str)
		}

		assert.Equal(t, expected, result)
	})
}
