package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	TestSorter(MergeSorter{}, t)
}
