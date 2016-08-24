package sorting

import "github.com/andreasstrack/datastructures"

// BubbleSorter is a type implementing the
// BubbleSort algorithm.
type BubbleSorter struct {
}

// Sort implements the BubbleSort algorithm.
func (b BubbleSorter) Sort(cl *datastructures.ComparableList) {
	for b.sortStep(cl) {
	}
}

// Name returns the name of the bubble sorter.
func (b BubbleSorter) Name() string {
	return "Bubble Sort"
}

func (b BubbleSorter) sortStep(cl *datastructures.ComparableList) (modifiedSomething bool) {
	modifiedSomething = false
	l := len(*cl)
	for i := 0; i < l-1; i++ {
		if (*cl)[i+1].LessThan((*cl)[i]) {
			modifiedSomething = true
			cl.Swap(i, i+1)
		}
	}

	return
}
