package stringsx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalize(t *testing.T) {
	// TODO: Реализовать тесты для Normalize
	t.Run("should normalize empty string", func(t *testing.T) {
		result := Normalize("")
		assert.Empty(t, result)
	})

	t.Run("should normalize simple string", func(t *testing.T) {
		result := Normalize("hello")
		assert.Equal(t, "hello", result)
	})

	t.Run("should normalize string with spaces", func(t *testing.T) {
		result := Normalize("  hello  world  ")
		assert.Equal(t, "hello world", result)
	})
}

func TestSplit(t *testing.T) {
	// TODO: Реализовать тесты для Split
	t.Run("should split empty string", func(t *testing.T) {
		result := Split("", ",")
		assert.Empty(t, result)
	})

	t.Run("should split string with separator", func(t *testing.T) {
		result := Split("a,b,c", ",")
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("should return single element when no separator found", func(t *testing.T) {
		result := Split("abc", ",")
		assert.Equal(t, []string{"abc"}, result)
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
