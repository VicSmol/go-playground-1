package setx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetXString(t *testing.T) {
	t.Run("should make setx string", func(t *testing.T) {
		var input *SetString = NewSetString()
		var expected *SetString = &SetString{set: make(map[string]struct{})}

		assert.Equal(t, expected, input)
	})

	t.Run("should add string to setx string", func(t *testing.T) {
		set := NewSetString()

		set.Add("1")

		assert.Equal(t, &SetString{set: map[string]struct{}{"1": {}}}, set)
	})

	t.Run("should remove string from setx string", func(t *testing.T) {
		set := NewSetString()

		set.Add("1")
		set.Remove("1")
		set.Remove("5")

		assert.Equal(t, &SetString{set: map[string]struct{}{}}, set)
	})

	t.Run("should check values contains in setx string", func(t *testing.T) {
		set := NewSetString()

		set.Add("1")

		assert.Equal(t, true, set.Contains("1"))
		assert.Equal(t, false, set.Contains("2"))
	})

	t.Run("should get slice from setx string", func(t *testing.T) {
		set := NewSetString()
		expected := []string{"1", "2", "3"}

		set.Add("1")
		set.Add("2")
		set.Add("3")

		assert.Equal(t, expected, set.ToSlice())
	})

	t.Run("should check that setx string is empty", func(t *testing.T) {
		set := NewSetString()
		emptySet := NewSetString()

		set.Add("1")

		assert.Equal(t, false, set.IsEmpty())
		assert.Equal(t, true, emptySet.IsEmpty())
	})
}
