package agstring

import (
	"regexp"
	"strings"
)

// ZeroWidth Constants
const (
	ZeroWidthSpace             = '\u200B'
	ZeroWidthSpaceNoBreakSpace = '\uFEFF'
	ZeroWidthJoiner            = '\u200D'
	ZeroWidthNonJoiner         = '\u200C'
)

// ReplaceWord replaces old strings into new only if olds occurs
// as a whole word according to regex word definition boundaries
func ReplaceWord(s, toBeReplaced string, replacements ...string) string {
	matcher := `\b(`
	for _, old := range replacements {
		matcher += regexp.QuoteMeta(old) + `|`
	}
	matcher = strings.TrimSuffix(matcher, `|`) + `)\b`
	return regexp.MustCompile(matcher).ReplaceAllLiteralString(s, toBeReplaced)
}

// ReplaceAll returns a string source with all elements of replacements replaced by toBeReplaced
func ReplaceAll(source string, toBeReplaced string, replacements ...string) string {
	for _, replacement := range replacements {
		source = strings.Replace(source, replacement, toBeReplaced, -1)
	}
	return source
}

// ReplaceZeroWidthChars returns a string with these characters:
// zero-width space:'\u200B', zero-width no break: '\uFEFF', zero-width joiner: '\u200D'
// and zero-width non-joiner: '\u200C' replaced by blank character
func ReplaceZeroWidthChars(s string) string {
	return ReplaceAll(s, "",
		string(ZeroWidthSpace),
		string(ZeroWidthSpaceNoBreakSpace),
		string(ZeroWidthJoiner),
		string(ZeroWidthNonJoiner))
}

// ReplaceMultispace replaces multiple spaces with one space and
// also trims space from both ends
func ReplaceMultispace(s string) string {
	stripper := regexp.MustCompile(`\s{2,}`)
	return strings.TrimSpace(stripper.ReplaceAllString(s, " "))
}

// ReplaceDayOrdinal replaces day ordinals (`st`, `nd`, `rd`, `th`)
// Default replaces with empty string.
func ReplaceDayOrdinal(s string, replacements ...string) string {
	var rep string
	if len(replacements) > 0 {
		rep = replacements[0]
	}
	return ReplaceAll(s, rep, "st", "nd", "th", "rd")
}

// ReplaceNewline replaces the newline character `\n`
// Default replaces with empty string.
func ReplaceNewline(s string, replacements ...string) string {
	var rep string
	if len(replacements) > 0 {
		rep = replacements[0]
	}
	return strings.Replace(s, "\n", rep, -1)
}
