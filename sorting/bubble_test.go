package sorting

import "testing"

func TestBubbleSorterOnRandomList(t *testing.T) {
	TestSorter(BubbleSorter{}, t)
}
