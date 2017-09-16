package algo

import (
	"github.com/andreasstrack/data/set"
	R "github.com/andreasstrack/util/reflect"
)

func HasDuplicates(sliceAsInterface interface{}) bool {
	var err error
	var slice *[]interface{}
	if slice, err = R.GetSlice(sliceAsInterface); err != nil {
		panic(err.Error())
	}
	var s *set.Set
	if s, err = set.NewSetFromSlice(sliceAsInterface); err != nil {
		panic(err.Error())
	}
	return len(*slice) > s.Size()
}
