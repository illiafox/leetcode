package hard

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		lists    [][]int // each inner slice is one list; empty slice => nil
		expected []int
	}{
		{nil, []int{}},
		{[][]int{{}, {}, {}}, []int{}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{{1, 1, 1}, {1, 1}, {1}}, []int{1, 1, 1, 1, 1, 1}},
		{[][]int{{-10, -1, 0}, {-5, -2}, {}}, []int{-10, -5, -2, -1, 0}},
		{[][]int{{5}, {3}, {4}, {1}, {2}}, []int{1, 2, 3, 4, 5}},
		{[][]int{{}, {1}, {0, 2, 100}, {3, 4, 5, 6, 7}, {-1}}, []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 100}},
	}

	for i, tt := range tests {
		var inputs []*ListNode
		for _, arr := range tt.lists {
			if len(arr) == 0 {
				inputs = append(inputs, nil)
			} else {
				inputs = append(inputs, listFrom(arr))
			}
		}

		got := sliceFromList(mergeKLists(inputs))
		if !reflect.DeepEqual(got, tt.expected) {
			if len(got)+len(tt.expected) == 0 { // nil or empty slice are equal for this test
				continue
			}
			t.Errorf("case %d: got %v; want %v", i, got, tt.expected)
		}
	}
}
