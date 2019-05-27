package agstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHead(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"one"}, "one"},
		{[]string{"one", "two"}, "one"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, Head(tt.input...))
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"one"}, "one"},
		{[]string{"one", "two"}, "two"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, Last(tt.input...))
	}
}

func TestTail(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{nil, nil},
		{[]string{}, nil},
		{[]string{"one"}, nil},
		{[]string{"one", "two", "three"}, []string{"two", "three"}},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, Tail(tt.input...))
	}
}

func TestNoLast(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{nil, nil},
		{[]string{}, nil},
		{[]string{"one"}, nil},
		{[]string{"one", "two", "three"}, []string{"one", "two"}},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, NoLast(tt.input...))
	}
}
