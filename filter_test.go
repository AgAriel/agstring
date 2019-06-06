package agstring

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		input    []string
		p        Predicate
		expected []string
	}{
		{
			input:    nil,
			p:        func(s string) bool { return true },
			expected: []string{},
		}, {
			input:    []string{},
			p:        func(s string) bool { return true },
			expected: []string{},
		}, {
			input:    []string{"one"},
			p:        func(s string) bool { return strings.Contains(s, "ne") },
			expected: []string{"one"},
		}, {
			input:    []string{"one", "two", "three"},
			p:        func(s string) bool { return strings.Contains(s, "e") },
			expected: []string{"one", "three"},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, Filter(tt.input, tt.p))
	}
}

func TestSliceFilter(t *testing.T) {
	tests := []struct {
		input    [][]string
		p        PredicateSlice
		expected [][]string
	}{
		{
			input:    nil,
			p:        func(ls []string) bool { return true },
			expected: [][]string{},
		}, {
			input:    [][]string{},
			p:        func(ls []string) bool { return true },
			expected: [][]string{},
		}, {
			input:    [][]string{{"one", "two", "three"}, {"four", "five", "six"}, {"seven", "eight", "nine"}},
			p:        func(ls []string) bool { return IsEqual(ls, []string{"four", "five", "six"}) },
			expected: [][]string{{"four", "five", "six"}},
		}, {
			input: [][]string{{"one", "two", "three"}, {"four", "five", "six"}, {"seven", "eight", "nine"}},
			p: func(ls []string) bool {
				return strings.Count(strings.Join(ls, " "), "e") >= 2
			},
			expected: [][]string{{"one", "two", "three"}, {"seven", "eight", "nine"}},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, FilterSlice(tt.input, tt.p))
	}
}
