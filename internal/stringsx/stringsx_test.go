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
		var input = []string{"", "a", "a,b", "a,b,c", "a,,b,c,", ",,a, ,b,,c,,"}
		var expected = [][]string{{}, {"a"}, {"a", "b"}, {"a", "b", "c"}, {"a", "b", "c"}, {"a", " ", "b", "c"}}
		var result = make([][]string, len(input))
		var separator = ","

		for index, input := range input {
			result[index] = Split(input, separator)
		}

		assert.Equal(t, expected, result)
	})
}

func TestJoin(t *testing.T) {
	// TODO: Реализовать тесты для Join
	t.Run("should join empty slice", func(t *testing.T) {
		result := Join([]string{}, ",")
		assert.Empty(t, result)
	})

	t.Run("should join slice with separator", func(t *testing.T) {
		result := Join([]string{"a", "b", "c"}, ",")
		assert.Equal(t, "a,b,c", result)
	})

	t.Run("should join slice with empty separator", func(t *testing.T) {
		result := Join([]string{"a", "b", "c"}, "")
		assert.Equal(t, "abc", result)
	})
}

func TestParseKV(t *testing.T) {
	// TODO: Реализовать тесты для ParseKV
	t.Run("should parse empty string", func(t *testing.T) {
		result := ParseKV("")
		assert.Empty(t, result)
	})

	t.Run("should parse simple key=value", func(t *testing.T) {
		result := ParseKV("key=value")
		assert.Equal(t, map[string]string{"key": "value"}, result)
	})

	t.Run("should parse multiple key=value pairs", func(t *testing.T) {
		result := ParseKV("key=value;key2=value2")
		assert.Equal(t, map[string]string{"key": "value", "key2": "value2"}, result)
	})

	t.Run("should handle empty values", func(t *testing.T) {
		result := ParseKV("key=;key2=value")
		assert.Equal(t, map[string]string{"key": "", "key2": "value"}, result)
	})
}
