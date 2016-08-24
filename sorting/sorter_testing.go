package sorting

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/andreasstrack/datastructures"
)

// TestSorter executes a test on a sorter and returns
// a bool indicating success along with a description of the
// working of the algorithm.
func TestSorter(s Sorter, t *testing.T) {
	success, description := TestSorterOnRandomList(s, 50, 1000)

	if success {
		fmt.Printf(description)
	} else {
		t.Error(description)
	}

	fmt.Printf("\n")
}

// TestSorterOnRandomList runs a Sorter on a randomly
// sorted list of size n. It returns a bool indicating success along
// with a description of the working of the algorithm.
func TestSorterOnRandomList(s Sorter, n, upperLimit int) (success bool, description string) {
	cl := datastructures.RandomIntList(n, upperLimit)
	testDescription := fmt.Sprintf("%s on random list(%d,%d)", s.Name(), n, upperLimit)
	return TestSorterOnList(s, cl, testDescription)
}

// TestSorterOnList tests a Sorter on a ComparableList and returns a bool
// indicating success along with a description of the working of the algorithm.
func TestSorterOnList(s Sorter, cl *datastructures.ComparableList, testDescription string) (success bool, description string) {
	var buf bytes.Buffer
	buf.WriteString(testDescription + ":\n")
	buf.WriteString(fmt.Sprintf("%s", *cl))
	buf.WriteString("\n->\n")
	s.Sort(cl)
	buf.WriteString(fmt.Sprintf("%s", *cl))
	buf.WriteString("\n")
	description = buf.String()
	success = cl.IsSorted()
	return
}
