package algorithm_test

import (
	"heap_sort/algorithm"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	res := algorithm.Sort([]int{3, 2, 1, 5, 4})
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}
