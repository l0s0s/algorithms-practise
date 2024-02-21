package algorithm_test

import (
	"insertion_sort/algorithm"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	expected := []int{1, 2, 3, 4, 5, 6}
	actual := algorithm.Sort(arr)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
