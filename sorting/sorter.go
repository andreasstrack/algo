package sorting

// A Sorter can sort a ComparableList.
type Sorter interface {
	// Sort sorts comparables in-place.
	Sort(comparables *data.ComparableList)

	// Name returns the name of the sorter.
	Name() string
}
