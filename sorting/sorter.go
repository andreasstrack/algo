package sorting

import "github.com/andreasstrack/datastructures"

// A Sorter can sort a ComparableList.
type Sorter interface {
	// Sort sorts comparables in-place.
	Sort(comparables *datastructures.ComparableList)

	// Name returns the name of the sorter.
	Name() string
}
