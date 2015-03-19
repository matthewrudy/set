package set

// Set is a set
type Set struct {
	m map[interface{}]bool
}

// New creates a new Set
func New() Set {
	s := Set{}
	s.Clear()
	return s
}

// Add add an element to the set
func (s *Set) Add(el interface{}) {
	s.m[el] = true
}

// Remove removes an element from the set
func (s *Set) Remove(el interface{}) {
	delete(s.m, el)
}

// Clear resets the set to empty
func (s *Set) Clear() {
	s.m = make(map[interface{}]bool)
}

// Size returns the cardinality of the set
func (s *Set) Size() int {
	return len(s.m)
}

// Has returns a boolean indicating whether the set contains el
func (s *Set) Has(el interface{}) bool {
	return s.m[el] == true
}
