package agstring

import "github.com/thoas/go-funk"

// Predicate is a function to filter slices with a string as input
type Predicate func(string) bool

// PredicateSlice is a function to filter slices with a slice as input
type PredicateSlice func([]string) bool

// Filter filters `ls` according to `p`
func Filter(ls []string, p Predicate) []string {
	return funk.FilterString(ls, p)
}

// FilterSlice filters `ls` according to `p`
func FilterSlice(ls [][]string, p PredicateSlice) [][]string {
	results := [][]string{}
	for _, i := range ls {
		if p(i) {
			results = append(results, i)
		}
	}
	return results
}
