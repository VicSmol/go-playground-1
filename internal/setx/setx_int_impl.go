package setx

// Реализовать set[T] на map[T]struct{} (пока без generics — как “SetString”, “SetInt”)

type SetInt struct {
	set map[int]struct{}
}

func NewSet() *SetInt {
	return &SetInt{set: make(map[int]struct{})}
}

func (s *SetInt) Add(value int) {
	s.set[value] = struct{}{}
}

func (s *SetInt) Remove(value int) {
	delete(s.set, value)
}

func (s *SetInt) Contains(value int) bool {
	_, ok := s.set[value]

	return ok
}

func (s *SetInt) Size() int {
	return len(s.set)
}

func (s *SetInt) ToSlice() []int {
	slice := make([]int, len(s.set))
	i := 0

	for value := range s.set {
		slice[i] = value
		i++
	}

	return slice
}

func (s *SetInt) IsEmpty() bool {
	return len(s.set) == 0
}
