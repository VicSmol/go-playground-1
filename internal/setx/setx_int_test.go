package setx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetXInt(t *testing.T) {
	t.Run("should make setx int", func(t *testing.T) {
		var input *SetInt = NewSet()
		var expected *SetInt = &SetInt{set: make(map[int]struct{})}

		assert.Equal(t, expected, input)
	})

	t.Run("should add int to setx int", func(t *testing.T) {
		set := NewSet()

		set.Add(1)

		assert.Equal(t, &SetInt{set: map[int]struct{}{1: {}}}, set)
	})

	t.Run("should remove int from setx int", func(t *testing.T) {
		set := NewSet()

		set.Add(1)
		set.Remove(1)
		set.Remove(5)

		assert.Equal(t, &SetInt{set: map[int]struct{}{}}, set)
	})

	t.Run("should check values contains in setx int", func(t *testing.T) {
		set := NewSet()

		set.Add(1)

		assert.Equal(t, true, set.Contains(1))
		assert.Equal(t, false, set.Contains(2))
	})

	t.Run("should get slice from setx int", func(t *testing.T) {
		set := NewSet()
		expected := []int{1, 2, 3}

		set.Add(1)
		set.Add(2)
		set.Add(3)

		assert.Equal(t, expected, set.ToSlice())
	})

	t.Run("should check that setx int is empty", func(t *testing.T) {
		set := NewSet()
		emptySet := NewSet()

		set.Add(1)

		assert.Equal(t, false, set.IsEmpty())
		assert.Equal(t, true, emptySet.IsEmpty())
	})
}
