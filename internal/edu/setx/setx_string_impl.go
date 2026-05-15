package setx

// Реализовать set[T] на map[T]struct{} (пока без generics — как “SetString”, “SetInt”)

type SetString struct {
	set map[string]struct{}
}

func NewSetString() *SetString {
	return &SetString{set: make(map[string]struct{})}
}

func (s *SetString) Add(value string) {
	s.set[value] = struct{}{}
}

func (s *SetString) Remove(value string) {
	delete(s.set, value)
}

func (s *SetString) Contains(value string) bool {
	_, ok := s.set[value]

	return ok
}

func (s *SetString) Size() int {
	return len(s.set)
}

func (s *SetString) ToSlice() []string {
	slice := make([]string, len(s.set))
	i := 0

	for value := range s.set {
		slice[i] = value
		i++
	}

	return slice
}

func (s *SetString) IsEmpty() bool {
	return len(s.set) == 0
}
