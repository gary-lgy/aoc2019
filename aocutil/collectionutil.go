package aocutil

// ReverseIntSlice reverses an integer slice
func ReverseIntSlice(slice []int) []int {
	s := make([]int, len(slice))
	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		s[i], s[j] = slice[j], slice[i]
	}
	return s
}

// IntSliceEqual checks if two integer slices contains the same elements
func IntSliceEqual(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}
