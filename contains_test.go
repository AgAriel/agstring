package agstring

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsSubString(t *testing.T) {
	testCases := []struct {
		holder   []string
		searched string
		expected bool
	}{
		{nil, "a", false},
		{[]string{"anna"}, "nn", true},
		{[]string{"anna", "banana"}, "na", true},
		{[]string{"anna", "banana"}, "anna", true},
		{[]string{"anna", "banana"}, "orange", false},
		{[]string{"", "banana"}, "orange", false},
		{[]string{"anna", "banana"}, "", true},
		{[]string{"a", "b"}, "a", true},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, ContainsSubString(tt.holder, tt.searched))
	}
}

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		holder, searched []string
		expected         bool
	}{
		{nil, []string{"a"}, false},
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a", "c"}, true},
		{[]string{"a", "b"}, []string{"d", "e"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, ContainsAny(tt.holder, tt.searched...))
	}
}

func TestStringContainsAny(t *testing.T) {
	testCases := []struct {
		holder   string
		searched []string
		expected bool
	}{
		{"", []string{"a"}, false},
		{"a", []string{"a"}, true},
		{"ab", []string{"a"}, true},
		{"ab", []string{"a", "c"}, true},
		{"ab", []string{"c", "d"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, StringContainsAny(tt.holder, tt.searched...))
	}
}
func TestCombinedContainsAny(t *testing.T) {
	testCases := []struct {
		holder, searched []string
		expected         bool
	}{
		{[]string{""}, []string{"a"}, false},
		{[]string{"", "a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"ab"}, true},
		{[]string{"a", "b"}, []string{"ba"}, false},
		{[]string{"ab"}, []string{"a", "c"}, true},
		{[]string{"ab"}, []string{"c", "d"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, CombinedContainsAny(tt.holder, tt.searched...))
	}
}
func TestSpaceCombinedContainsAny(t *testing.T) {
	testCases := []struct {
		holder, searched []string
		expected         bool
	}{
		{[]string{""}, []string{"a"}, false},
		{[]string{"", "a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"ab"}, false},
		{[]string{"a", "b"}, []string{"a b"}, true},
		{[]string{"a", "b"}, []string{"ba"}, false},
		{[]string{"ab"}, []string{"a", "c"}, true},
		{[]string{"ab"}, []string{"c", "d"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, SpaceCombinedContainsAny(tt.holder, tt.searched...))
	}
}

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		holder, searched []string
		expected         bool
	}{
		{nil, nil, true},
		{[]string{}, []string{}, true},
		{nil, []string{"a"}, false},
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a", "c"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, ContainsAll(tt.holder, tt.searched...))
	}
}

func TestStringContainsAll(t *testing.T) {
	testCases := []struct {
		holder   string
		searched []string
		expected bool
	}{
		{"", nil, true},
		{"", []string{}, true},
		{"", []string{"a"}, false},
		{"a", []string{"a"}, true},
		{"ab", []string{"a"}, true},
		{"ab", []string{"a", "c"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, StringContainsAll(tt.holder, tt.searched...))
	}
}

func TestCombinedContainsAll(t *testing.T) {
	testCases := []struct {
		holder, searched []string
		expected         bool
	}{
		{[]string{""}, []string{"a"}, false},
		{[]string{"", "a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a", "b"}, true},
		{[]string{"a", "b"}, []string{"a", "b", "c"}, false},
		{[]string{"a", "b"}, []string{"ab"}, true},
		{[]string{"a", "b"}, []string{"ba"}, false},
		{[]string{"ab"}, []string{"a", "c"}, false},
		{[]string{"ab"}, []string{"c", "d"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, CombinedContainsAll(tt.holder, tt.searched...))
	}
}

func TestSpaceCombinedContainsAll(t *testing.T) {
	testCases := []struct {
		holder, searched []string
		expected         bool
	}{
		{[]string{""}, []string{"a"}, false},
		{[]string{"", "a"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a"}, true},
		{[]string{"a", "b"}, []string{"a", "b"}, true},
		{[]string{"a", "b"}, []string{"a", "b", "c"}, false},
		{[]string{"a", "b"}, []string{"ab"}, false},
		{[]string{"a", "b"}, []string{"a b"}, true},
		{[]string{"a", "b"}, []string{"ba"}, false},
		{[]string{"ab"}, []string{"a", "c"}, false},
		{[]string{"ab"}, []string{"c", "d"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, SpaceCombinedContainsAll(tt.holder, tt.searched...))
	}
}

func TestIndexContainingSubString(t *testing.T) {
	testCases := []struct {
		holder   []string
		searched string
		expected int
	}{
		{nil, "a", -1},
		{[]string{"anna"}, "nn", 0},
		{[]string{"anna", "banana"}, "na", 0},
		{[]string{"anna", "banana"}, "ana", 1},
		{[]string{"anna", "banana"}, "orange", -1},
		{[]string{"", "banana"}, "orange", -1},
		{[]string{"anna", "banana"}, "", 0},
		{[]string{"a", "b"}, "a", 0},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, IndexContainingSubString(tt.holder, tt.searched))
	}
}

func TestStringIndexContainingSubString(t *testing.T) {
	testCases := []struct {
		holder   string
		searched []string
		expected int
	}{
		{"a", nil, -1},
		{"banana", []string{"steve", "anna"}, -1},
		{"banana", []string{"steve", "ana"}, 1},
		{"orange", []string{"anna", "banana"}, -1},
		{"orange", []string{"", "banana"}, 0},
		{"", []string{"anna", "banana"}, -1},
		{"a", []string{"a", "b"}, 0},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, StringIndexContainingSubString(tt.holder, tt.searched...))
	}
}

func TestContainsSlice(t *testing.T) {
	trim1 := func(s string) string { return strings.TrimPrefix(s, "1") }
	trim2 := func(s string) string { return strings.TrimPrefix(s, "2") }

	trueTestCases := []struct {
		slice      [][]string
		s          []string
		transforms []Transform
	}{
		{
			slice: [][]string{{"1", "2", "3"}, {"6", "7", "8"}, {"12", "13", "14"}},
			s:     []string{"1", "2", "3"},
		},
		{
			slice: [][]string{{"1", "2", "3"}, {}, {"12", "13", "14"}},
			s:     []string{},
		},
		{
			slice:      [][]string{{"6", "7", "8"}, {"1", "2", "3"}},
			s:          []string{"11", "12", "13"},
			transforms: []Transform{trim1},
		},
		{
			slice:      [][]string{{"6", "7", "8"}, {"11", "12", "13"}, {"1", "2", "3"}},
			s:          []string{"211", "212", "213"},
			transforms: []Transform{trim2, trim1},
		},
	}

	for _, tt := range trueTestCases {
		require.True(t, SliceContains(tt.slice, tt.s, tt.transforms...))
	}

	falseTestCases := []struct {
		slice      [][]string
		s          []string
		transforms []Transform
	}{
		{
			slice: [][]string{{"1", "2", "3"}, {"6", "7", "8"}, {"12", "13", "14"}},
			s:     []string{"16", "26", "36"},
		},
		{
			slice:      [][]string{{"6", "7", "8"}, {"61", "62", "63"}},
			s:          []string{"1", "2", "3"},
			transforms: []Transform{trim1},
		},
		{
			slice:      [][]string{{"6", "7", "8"}, {"61", "62", "63"}},
			s:          []string{"1", "2", "3"},
			transforms: []Transform{trim2, trim1},
		},
	}

	for _, tt := range falseTestCases {
		require.False(t, SliceContains(tt.slice, tt.s, tt.transforms...))
	}
}
