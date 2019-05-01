package agstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEqual(t *testing.T) {
	addA := func(s string) string { return s + "a" }
	testCases := []struct {
		input, probe []string
		transforms   []Transform
		expected     bool
	}{
		{input: []string{}, probe: []string{}, expected: true},
		{input: []string{}, probe: nil, expected: true},
		{input: nil, probe: nil, expected: true},
		{input: []string{"", "a"}, probe: []string{"a", ""}, expected: false},
		{input: []string{}, probe: []string{},
			transforms: []Transform{UnitTransform}, expected: true},
		{input: []string{"a"}, probe: []string{"a"},
			transforms: []Transform{UnitTransform}, expected: true},
		{input: []string{"aa"}, probe: []string{"a"},
			transforms: []Transform{addA, UnitTransform}, expected: true},
		{input: []string{"ab"}, probe: []string{"a"},
			transforms: []Transform{addA, UnitTransform}, expected: false},
		{input: []string{"ab", "aa"}, probe: []string{"a", "a"},
			transforms: []Transform{addA, UnitTransform}, expected: false},
	}

	for _, testCase := range testCases {
		require.Equalf(t, testCase.expected,
			IsEqual(testCase.input, testCase.probe, testCase.transforms...),
			"input is %v", testCase)
	}
}

func TestMustEqual(t *testing.T) {
	addA := func(s string) string { return s + "a" }
	testCases := []struct {
		input, probe []string
		transforms   []Transform
		expected     error
	}{
		{input: []string{}, probe: []string{}},
		{input: []string{}, probe: nil},
		{input: nil, probe: nil},
		{input: []string{"a", "a"}, probe: []string{"a", "a"}},
		{input: []string{}, probe: []string{},
			transforms: []Transform{UnitTransform}},
		{input: []string{"a"}, probe: []string{"a"},
			transforms: []Transform{UnitTransform}},
		{input: []string{"aa"}, probe: []string{"a"},
			transforms: []Transform{addA, UnitTransform}},
	}

	for _, testCase := range testCases {
		require.Equalf(t, testCase.expected,
			MustEqual(testCase.input, testCase.probe, testCase.transforms...),
			"inpur is %v", testCase)
	}
}
