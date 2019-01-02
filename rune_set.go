package set

// RuneSet rune set.
type RuneSet struct {
	list map[rune]struct{}
}

// Has returns true if set contains v value.
func (s *RuneSet) Has(v rune) bool {
	_, ok := s.list[v]
	return ok
}

// Add adds a value to the set.
func (s *RuneSet) Add(v rune) {
	s.list[v] = struct{}{}
}

// Remove removes a value from the set.
func (s *RuneSet) Remove(v rune) {
	delete(s.list, v)
}

// Clear clears the whole list.
func (s *RuneSet) Clear() {
	s.list = make(map[rune]struct{})
}

// Size returns the size of the set
func (s *RuneSet) Size() int {
	return len(s.list)
}

// NewRuneSet creates a new rune set.
func NewRuneSet() *RuneSet {
	s := &RuneSet{}
	s.list = make(map[rune]struct{})
	return s
}

// AddMulti Add multiple values in the RuneSet.
func (s *RuneSet) AddMulti(list ...rune) {
	for _, v := range list {
		s.Add(v)
	}
}

// RuneFilterFunc defines a function that filters out certain values.
type RuneFilterFunc func(v rune) bool

// Filter returns a subset, that contains only the values that satisfy the given predicate P
func (s *RuneSet) Filter(P RuneFilterFunc) *RuneSet {
	res := NewRuneSet()
	for v := range s.list {
		if P(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Union returns the union of the current and another rune set.
func (s *RuneSet) Union(s2 *RuneSet) *RuneSet {
	res := NewRuneSet()
	for v := range s.list {
		res.Add(v)
	}

	for v := range s2.list {
		res.Add(v)
	}
	return res
}

// Intersect returns the intersect of the current and another rune set.
func (s *RuneSet) Intersect(s2 *RuneSet) *RuneSet {
	res := NewRuneSet()
	for v := range s.list {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the sub from s, that doesn't exists in s2 (param)
func (s *RuneSet) Difference(s2 *RuneSet) *RuneSet {
	res := NewRuneSet()
	for v := range s.list {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}
