package aocutil

// BitSet is a set of bits
type BitSet uint32

// Has tests whether the bit at flag is true
func (s BitSet) Has(flag int) bool {
	return s&(1<<flag) != 0
}

// Set sets the bit at flag to true
func (s *BitSet) Set(flag int) {
	*s |= 1 << flag
}

// Unset sets the bit at flag to false
func (s *BitSet) Unset(flag int) {
	*s &= ^(1 << flag)
}
