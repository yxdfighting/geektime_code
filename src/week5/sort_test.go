package week5

import (
	"testing"
	"week5/sortAlgol"
)

func TestSort(t *testing.T) {
	nums := []int{4, 5, 6, 2, 3, 9, 3, 2, 1, 5}
	t.Logf("input:%v\n", nums)
	sortAlgol.HeapSort(nums)
	t.Logf("sorted:%v\n", nums)
}
