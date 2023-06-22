package stringfns

// HasPrefix returns true if s starts with prefix, false otherwise
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix returns true if s ends with prefix, false otherwise
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains returns true if s contains substr in some part of it, false otherwise
func Contains(s, substr string) bool {
	for i := 0; i < len(s) - len(substr) + 1; i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
