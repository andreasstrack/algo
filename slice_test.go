package algo

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestHasDuplicatesNoDuplicates(t *testing.T) {
	tt := T.NewT(t)
	var slice []interface{}
	slice = append(slice, 1, 2, "2", 3.1)
	tt.Assert(!HasDuplicates(&slice), "Check: no duplicates in %#v", slice)
}

func TestHasDuplicatesDuplicates(t *testing.T) {
	tt := T.NewT(t)
	var slice []interface{}
	slice = append(slice, 1, "2", 3.1, 3.1)
	tt.Assert(HasDuplicates(&slice), "Check: have duplicates in %#v", slice)
}
