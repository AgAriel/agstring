package agstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplaceAll(t *testing.T) {
	testCasesWithASlice := []struct {
		input        string
		toBeReplaced string
		replacements []string
		expected     string
	}{
		{"January20190301June", "j", []string{"J"},
			"january20190301june"},
		{"My saster het a heth pet which as", "", []string{"et", "as"},
			"My ster h a hh p which "},
	}

	for _, tt := range testCasesWithASlice {
		require.Equal(t, tt.expected, ReplaceAll(tt.input, tt.toBeReplaced, tt.replacements...))
	}
}

func TestZeroWidthSpace(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "Test:\u200B:\u200B:Test", expected: "Test:::Test"},
		{input: "Test:\u200C:\u200B:Test", expected: "Test:::Test"},
		{input: "Test:\u200D:\u200B:Test", expected: "Test:::Test"},
		{input: "Test:\uFEFF:\u200B:Test", expected: "Test:::Test"},
	}

	for _, testCase := range testCases {
		require.Equal(t, ReplaceZeroWidthChars(testCase.input), testCase.expected)
	}
}

func TestReplaceMultispace(t *testing.T) {
	tests := []string{"   a   b   c d    ef  ", "     123   ", "1\n\n123"}
	expects := []string{"a b c d ef", "123", "1 123"}

	for i, tt := range tests {
		require.Equal(t, ReplaceMultispace(tt), expects[i])
	}
}

func TestReplaceWord(t *testing.T) {
	tests := []struct {
		s           string
		replacement string
		olds        []string
		expected    string
	}{
		{
			"port saint jose",
			"nntt",
			[]string{"nt"},
			"port saint jose",
		}, {
			"po saint jose",
			"port",
			[]string{"po"},
			"port saint jose",
		}, {
			"port st jose",
			"saint",
			[]string{"st"},
			"port saint jose",
		}, {
			"port saint hose",
			"jose",
			[]string{"hose"},
			"port saint jose",
		}, {"port-saint-hose",
			"jose",
			[]string{"hose"},
			"port-saint-jose",
		}, {
			"port-saint/hose",
			"jose",
			[]string{"hose"},
			"port-saint/jose",
		}, {"port-saint/h(os)e",
			"jose",
			[]string{"(os)"},
			"port-saint/hjosee",
		}, {
			"po saint jose re po",
			"port",
			[]string{"po", "re"},
			"port saint jose port port",
		},
	}
	for _, tt := range tests {
		actual := ReplaceWord(tt.s, tt.replacement, tt.olds...)
		require.Equal(t, tt.expected, actual)
	}
}

func TestDayOrdinalReplacer(t *testing.T) {
	testCases := []struct {
		input       string
		replaceWith string
		expected    string
	}{
		{
			"2nd Jan",
			"",
			"2 Jan",
		}, {
			"3rd Jan",
			"XX",
			"3XX Jan",
		}, {
			"4th Jan",
			"ZZZ",
			"4ZZZ Jan",
		},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, ReplaceDayOrdinal(tt.input, tt.replaceWith))
	}

	require.Equal(t, "1 Jan", ReplaceDayOrdinal("1st Jan"))
}

func TestReplaceNewline(t *testing.T) {
	input := "Hi\nThere"
	testCases := []struct {
		replaceWith string
		expected    string
	}{
		{
			"",
			"HiThere",
		}, {
			" ",
			"Hi There",
		}, {
			"XXX",
			"HiXXXThere",
		},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.expected, ReplaceNewline(input, tt.replaceWith))
	}

	require.Equal(t, "HiThere", ReplaceNewline(input))
}
