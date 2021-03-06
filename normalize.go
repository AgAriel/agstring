package agstring

import (
	"strings"
)

// Normalize first replace ZeroWidth chars then trims it and finally lowercase it
func Normalize(s string) string {
	return strings.ToLower(TrimAndRemoveZeroWidthChars(s))
}

// NormalizeDiacritics remove diacritics from a normalized string
func NormalizeDiacritics(s string) string {
	return Normalize(RemoveDiacritics(s))
}

// NormalizeDiacriticsAndNonAlnum first remove diacritics from a normalized string then remove
// all non alphanumeric characters including whitespaces in the middle
func NormalizeDiacriticsAndNonAlnum(s string) string {
	return RemoveNonAlnum(NormalizeDiacritics(s))
}
