package agstring

// Head returns the first element of given list or empty string when the list is empty.
func Head(ls ...string) string { return Nth(ls, 0) }

// Last returns the last element of given list or empty string when the list is empty.
func Last(ls ...string) string { return Nth(ls, len(ls)-1) }

// Nth returns nth element of given slice or empty string if out of limits
func Nth(ls []string, n int) string {
	if len(ls) == 0 || n < 0 || n >= len(ls) {
		return ""
	}
	return ls[n]
}

// Tail returns a slice without the first element, otherwise nil
func Tail(ls ...string) []string {
	if len(ls) <= 1 {
		return nil
	}
	return ls[1:]
}

// NoLast returns a slice without the last element, otherwise nil
func NoLast(ls ...string) []string {
	if len(ls) <= 1 {
		return nil
	}
	return ls[:len(ls)-1]
}
