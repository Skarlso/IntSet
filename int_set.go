package set

// IntSet integer set.
type IntSet struct {
	list map[int]struct{} //empty structs occupy 0 memory
}

// Has returns true if set contains v value.
func (s *IntSet) Has(v int) bool {
	_, ok := s.list[v]
	return ok
}

// Add adds a value to the set.
func (s *IntSet) Add(v int) {
	s.list[v] = struct{}{}
}

// Remove removes a value from the set.
func (s *IntSet) Remove(v int) {
	delete(s.list, v)
}

// Clear clears the whole list.
func (s *IntSet) Clear() {
	s.list = make(map[int]struct{})
}

// Size returns the size of the set
func (s *IntSet) Size() int {
	return len(s.list)
}

// NewIntSet creates a new integer set.
func NewIntSet() *IntSet {
	s := &IntSet{}
	s.list = make(map[int]struct{})
	return s
}

// AddMulti Add multiple values in the IntSet.
func (s *IntSet) AddMulti(list ...int) {
	for _, v := range list {
		s.Add(v)
	}
}

// IntFilterFunc defines a function that filters out certain values.
type IntFilterFunc func(v int) bool

// Filter returns a subset, that contains only the values that satisfy the given predicate P
func (s *IntSet) Filter(P IntFilterFunc) *IntSet {
	res := NewIntSet()
	for v := range s.list {
		if P(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Union returns the union of the current and another integer set.
func (s *IntSet) Union(s2 *IntSet) *IntSet {
	res := NewIntSet()
	for v := range s.list {
		res.Add(v)
	}

	for v := range s2.list {
		res.Add(v)
	}
	return res
}

// Intersect returns the intersect of the current and another integer set.
func (s *IntSet) Intersect(s2 *IntSet) *IntSet {
	res := NewIntSet()
	for v := range s.list {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the subset from s that doesn't exists in s2.
func (s *IntSet) Difference(s2 *IntSet) *IntSet {
	res := NewIntSet()
	for v := range s.list {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}
