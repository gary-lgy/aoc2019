package aocutil

import "strconv"

// Reverse reverses str
func Reverse(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; j < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

// Digits converts each character in digits into a digit represented as an int
func Digits(str string) ([]int, error) {
	var digits []int
	for _, ch := range str {
		digit, err := strconv.ParseInt(string(ch), 10, 32)
		if err != nil {
			return nil, err
		}
		digits = append(digits, int(digit))
	}
	return digits, nil
}
