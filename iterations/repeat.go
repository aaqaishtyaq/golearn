package iterations

// Repeat returns a new string consisting of count copies of the string s.
// It panics if count is negative or if the result of (len(s) * count) overflows.
func Repeat(char string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += char
	}
	return repeated
}
