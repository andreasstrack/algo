package sorting

// MergeSorter is a type implementing the
// merge sort algorithm.
type MergeSorter struct {
}

// Sort implements the merge sort algorithm.
func (ms MergeSorter) Sort(cl *data.ComparableList) {
	ms.sortIter(cl)
}

// Name returns the name of the merge sort algorithm.
func (ms MergeSorter) Name() string {
	return "Merge Sort"
}

func (ms MergeSorter) sortIter(cl *data.ComparableList) {
	if cl == nil {
		return
	} else if len(*cl) == 1 {
		return
	}

	cll, clr := cl.SplitToHalves()
	ms.sortIter(cll)
	ms.sortIter(clr)
	*cl = *ms.merge(cll, clr)
}

func (ms MergeSorter) merge(cll, clr *data.ComparableList) *data.ComparableList {
	var result data.ComparableList
	var il, ir int
	ll := len(*cll)
	lr := len(*clr)
	for il < ll || ir < lr {
		if il >= ll {
			result = append(result, (*clr)[ir])
			ir++
		} else if ir >= lr {
			result = append(result, (*cll)[il])
			il++
		} else if (*cll)[il].LessThan((*clr)[ir]) {
			result = append(result, (*cll)[il])
			il++
		} else {
			result = append(result, (*clr)[ir])
			ir++
		}
	}
	return &result
}
