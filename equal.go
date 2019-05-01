package agstring

import (
	"github.com/pkg/errors"
	funk "github.com/thoas/go-funk"
)

// IsEqual checks if two slices are equal after applying transforms on 2nd slice
func IsEqual(source, other []string, transforms ...Transform) bool {
	return (source == nil && other == nil) ||
		funk.Equal(source, Map(other, transforms...))
}

// MustEqual checks if two slices are equal after applying transforms on 2nd slice otherwise return error
func MustEqual(source, other []string, transforms ...Transform) error {
	if !IsEqual(source, other, transforms...) {
		return errors.Errorf("must be equal: expected %v, got %v", source, other)
	}
	return nil
}
